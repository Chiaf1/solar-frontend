package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
)

type EnergyAPIRepository struct {
	clinet  *http.Client
	baseURL string
}

func NewEnergyAPIRepository(baseURL string) *EnergyAPIRepository {
	return &EnergyAPIRepository{
		clinet: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (r *EnergyAPIRepository) GetToday() (models.ChartData, error) {
	url := fmt.Sprintf("%s/energy/today", r.baseURL)

	// 1. Create the request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.ChartData{}, err
	}

	// 2. Execute the request
	resp, err := r.clinet.Do(req)
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

func mapAPIResponseToChartData(apiResp []EnergyAPIPoint) models.ChartData
