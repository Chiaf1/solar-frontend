package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OpenMeteoRepository struct {
	client  *http.Client
	baseURL string
}

type Result struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
	} `json:"current"`
}

func NewOpenMeteoRepository(baseURL string) *OpenMeteoRepository {
	return &OpenMeteoRepository{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
	}
}

// Return current temperatura value
func (r *OpenMeteoRepository) GetCurrentTemperature(lat, lon float64) (float64, error) {
	url := fmt.Sprintf("%s?latitude=%f&longitude=%f&current=temperature_2m", r.baseURL, lat, lon)

	resp, err := r.client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var apiRest Result
	if err := json.NewDecoder(resp.Body).Decode(&apiRest); err != nil {
		return 0, err
	}

	return apiRest.Current.Temperature, nil
}
