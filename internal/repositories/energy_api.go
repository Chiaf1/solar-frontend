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

// Convert time stamp string into label format
func formatTimeStamp(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return ts
	}
	// Load local time zone (Italy)
	loc, err := time.LoadLocation("Europe/Rome")
	if err != nil {
		return t.Format("15:04")
	}
	// Conv from UTC to Europe/Rome
	return t.In(loc).Format("15:04")
}

func (r *EnergyAPIRepository) GetHistory() (map[string]models.ChartData, error) {
	return nil, nil
}

func (r *EnergyAPIRepository) GetKPI() (models.KPIData, error) {
	return models.KPIData{}, nil
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
