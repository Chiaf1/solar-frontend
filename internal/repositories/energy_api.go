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
	return r.GetTodayDaily()
}

// Retrive yesterday chart data from api
func (r *EnergyAPIRepository) GetYesterday() (models.ChartData, error) {
	url := fmt.Sprintf("%s/energy/yesterday", r.baseURL)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}

	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}
	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return emptyDailyChart(10 * time.Minute), fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// 4. Decode JSON
	var apiResp []EnergyAPIPoint
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return emptyDailyChart(10 * time.Minute), err
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

// Currently not used
func (r *EnergyAPIRepository) GetKPI() (models.KPIData, error) {
	return models.KPIData{}, nil
}

// Retrive today chart data from api with today endpoint
func (r *EnergyAPIRepository) getTodayOg() (models.ChartData, error) {
	url := fmt.Sprintf("%s/energy/today", r.baseURL)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}

	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}
	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return emptyDailyChart(10 * time.Minute), fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// 4. Decode JSON
	var apiResp []EnergyAPIPoint
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}

	// 5. Mapping to ChartData
	return mapAPIResponseToChartData(apiResp), err
}

// Retrive today chart data from api with daily endpoint
func (r *EnergyAPIRepository) GetTodayDaily() (models.ChartData, error) {
	window := "4m"
	loc, _ := time.LoadLocation("Europe/Rome")
	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	todayStr := today.Format("2006-01-02")
	url := fmt.Sprintf("%s/energy/daily?from=%s&to=%s&window=%s", r.baseURL, todayStr, todayStr, window)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}

	// 2. Execute the request
	resp, err := r.client.Do(req)
	if err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}
	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		return emptyDailyChart(10 * time.Minute), fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// 4. Decode JSON
	var apiResp []EnergyAPIDaily
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return emptyDailyChart(10 * time.Minute), err
	}

	// 5. Mapping to ChartData
	return mapAPIResponseToChartData(apiResp[0].Points), err
}
