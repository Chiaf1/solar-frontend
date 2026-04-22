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
	labels := make([]string, 0, len(apiResp))
	production := make([]float64, 0, len(apiResp))
	consumption := make([]float64, 0, len(apiResp))

	for _, point := range apiResp {
		labels = append(labels, formatTimeStamp(point.Time))
		production = append(production, point.Production)
		consumption = append(consumption, point.Consumption)
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
	return t.Format("15:04")
}

func (r *EnergyAPIRepository) GetHistory() (map[string]models.ChartData, error) {
	return nil, nil
}

func (r *EnergyAPIRepository) GetKPI() (models.KPIData, error) {
	return models.KPIData{}, nil
}
