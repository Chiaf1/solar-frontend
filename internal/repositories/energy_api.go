package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
)

type EnergyAPIRepository struct {
	client  *http.Client
	baseURL string
}

func NewEnergyAPIRepository(baseURL string) *EnergyAPIRepository {
	return &EnergyAPIRepository{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// Retrive today chart data from api
func (r *EnergyAPIRepository) GetToday() (models.ChartData, error) {
	url := fmt.Sprintf("%s/energy/today", r.baseURL)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.ChartData{}, err
	}

	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return models.ChartData{}, err
	}
	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return models.ChartData{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// 4. Decode JSON
	var apiResp []EnergyAPIPoint
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return models.ChartData{}, err
	}

	// 5. Mapping to ChartData
	return mapAPIResponseToChartData(apiResp), err
}

// Retrive yesterday chart data from api
func (r *EnergyAPIRepository) GetYesterday() (models.ChartData, error) {
	url := fmt.Sprintf("%s/energy/yesterday", r.baseURL)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.ChartData{}, err
	}

	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return models.ChartData{}, err
	}
	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return models.ChartData{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// 4. Decode JSON
	var apiResp []EnergyAPIPoint
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return models.ChartData{}, err
	}

	// 5. Mapping to ChartData
	return mapAPIResponseToChartData(apiResp), err
}

func (r *EnergyAPIRepository) GetHistory() (map[string]models.ChartData, error) {
	charts := map[string]models.ChartData{
		"chart-yesterday": emptyDailyChart(10 * time.Minute),
		"chart-minus-2":   emptyDailyChart(10 * time.Minute),
		"chart-minus-3":   emptyDailyChart(10 * time.Minute),
		"chart-minus-4":   emptyDailyChart(10 * time.Minute),
		"chart-minus-5":   emptyDailyChart(10 * time.Minute),
		"chart-minus-6":   emptyDailyChart(10 * time.Minute),
	}
	var err error
	// Yesterday
	charts["chart-yesterday"], err = r.GetYesterday()

	// Prev week (-2, ..., -6)
	// Url building
	window := "10m"
	loc, _ := time.LoadLocation("Europe/Rome")
	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)

	from := today.AddDate(0, 0, -6)
	to := today.AddDate(0, 0, -2)

	fromStr := from.Format("2006-01-02")
	toStr := to.Format("2006-01-02")

	url := fmt.Sprintf("%s/energy/daily?from=%s&to=%s&window=%s", r.baseURL, fromStr, toStr, window)
	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return charts, err
	}
	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return charts, err
	}
	defer resp.Body.Close()
	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return charts, fmt.Errorf("API returned status %d", resp.StatusCode)
	}
	// 4. Decode JSON
	var apiResp []EnergyAPIDaily
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return charts, err
	}

	// Data formatting for chart data map
	for _, day := range apiResp {
		chartData := mapAPIResponseToChartData(day.Points)

		key := ""
		switch day.Day {
		case from.AddDate(0, 0, 0).Format("2006-01-02"): // from + 0
			key = "chart-minus-6"
		case from.AddDate(0, 0, 1).Format("2006-01-02"): // from + 1
			key = "chart-minus-5"
		case from.AddDate(0, 0, 2).Format("2006-01-02"):
			key = "chart-minus-4"
		case from.AddDate(0, 0, 3).Format("2006-01-02"):
			key = "chart-minus-3"
		case from.AddDate(0, 0, 4).Format("2006-01-02"):
			key = "chart-minus-2"
		}

		if key != "" {
			charts[key] = chartData
		}
	}

	return charts, err
}

func (r *EnergyAPIRepository) GetKPI() (models.KPIData, error) {
	return models.KPIData{}, nil
}

// Convert EnergyAPIPoint slice into chartdata
func mapAPIResponseToChartData(apiResp []EnergyAPIPoint) models.ChartData {
	loc, _ := time.LoadLocation("Europe/Rome")

	step := detectStep(apiResp, 5*time.Minute, loc)

	timeline := buildDailyTimeline(step, loc)
	index := indexAPIPoints(apiResp, step, loc)

	labels := make([]string, 0, len(timeline))
	production := make([]*float64, 0, len(timeline))
	consumption := make([]*float64, 0, len(timeline))

	for _, t := range timeline {
		labels = append(labels, t.Format("15:04"))

		if point, ok := index[t]; ok {
			production = append(production, &point.Production)
			consumption = append(consumption, &point.Consumption)
		} else {
			production = append(production, nil)
			consumption = append(consumption, nil)
		}

	}

	return models.ChartData{
		Labels:      labels,
		Production:  production,
		Consumption: consumption,
	}
}

// Build daily timeline returns the labels for a comlpete day with a set step
func buildDailyTimeline(step time.Duration, loc *time.Location) []time.Time {
	start := time.Date(0, 1, 1, 0, 0, 0, 0, loc)
	end := start.Add(24 * time.Hour)

	var timeline []time.Time
	for t := start; !t.After(end); t = t.Add(step) {
		timeline = append(timeline, t)
	}
	return timeline
}

// Creates a map for all energy points with the time stamp as key
func indexAPIPoints(apiPoints []EnergyAPIPoint, step time.Duration, loc *time.Location) map[time.Time]EnergyAPIPoint {
	index := make(map[time.Time]EnergyAPIPoint)

	for _, p := range apiPoints {
		t, err := time.Parse(time.RFC3339, p.Time)
		if err != nil {
			continue
		}

		bucket := normalizeToBacket(t, step, loc)
		index[bucket] = p
	}
	return index
}

// Detect step duration in a slice of EnergyAPIPoint
func detectStep(apiResp []EnergyAPIPoint, fallback time.Duration, loc *time.Location) time.Duration {
	if len(apiResp) < 2 {
		return fallback
	}

	t1, err1 := time.Parse(time.RFC3339, apiResp[0].Time)
	t2, err2 := time.Parse(time.RFC3339, apiResp[1].Time)

	if err1 != nil || err2 != nil {
		return fallback
	}

	// Conv to local time before step calculation
	t1 = t1.In(loc)
	t2 = t2.In(loc)

	step := t2.Sub(t1)

	// Check step value
	if step <= 0 || step > time.Hour {
		return fallback
	}

	return step
}

// Normalize timestamp to time stamp starting from midnight with fixed step
func normalizeToBacket(ts time.Time, step time.Duration, loc *time.Location) time.Time {
	local := ts.In(loc)

	// Midnight in local time of the smae day
	startOfDay := time.Date(0, 1, 1, 0, 0, 0, 0, loc)

	elapsed := time.Duration(local.Hour())*time.Hour +
		time.Duration(local.Minute())*time.Minute +
		time.Duration(local.Second())*time.Second +
		time.Duration(local.Nanosecond())

	bucketIndex := elapsed / step

	return startOfDay.Add(bucketIndex * step)
}

// Helper function returns empty chart data
func emptyDailyChart(step time.Duration) models.ChartData {
	loc, _ := time.LoadLocation("Europe/Rome")
	timeline := buildDailyTimeline(step, loc)

	labels := make([]string, 0, len(timeline))
	production := make([]*float64, 0, len(timeline))
	consumption := make([]*float64, 0, len(timeline))

	for _, t := range timeline {
		labels = append(labels, t.Format("15:04"))
		production = append(production, nil)
		consumption = append(consumption, nil)
	}

	return models.ChartData{
		Labels:      labels,
		Production:  production,
		Consumption: consumption,
	}
}
