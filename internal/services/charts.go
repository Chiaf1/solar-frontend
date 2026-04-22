package services

import (
	"github.com/chiaf1/solar-frontend/internal/models"
	"github.com/chiaf1/solar-frontend/internal/repositories"
)

type ChartService struct {
	repo repositories.EnergyRepository
}

func NewChartService(repo repositories.EnergyRepository) *ChartService {
	return &ChartService{
		repo: repo,
	}
}

// Returns chart data for today's chart
func (s *ChartService) GetTodayChart() (models.ChartData, error) {
	data, err := s.repo.GetToday()
	if err != nil {
		return models.ChartData{}, err
	}
	return data, nil
}

// Returns chart data for history's charts
func (s *ChartService) GetHistoryCharts() (map[string]models.ChartData, error) {
	hystoryData := map[string]models.ChartData{
		"chart-yesterday": {
			Labels:      []string{"00:00", "04:00", "08:00", "12:00", "16:00", "20:00"},
			Production:  []float64{0.5, 1.2, 2.3, 1.8, 0.9, 0.3},
			Consumption: []float64{0.5, 1.2, 1.5, 1.0, 1.5, 0.3},
		},
		"chart-minus-2": {
			Labels:      []string{"00:00", "06:00", "12:00", "18:00"},
			Production:  []float64{0.2, 1.1, 2.0, 0.8},
			Consumption: []float64{0.2, 5.0, 3.0, 0.8},
		},
		"chart-minus-3": {
			Labels:      []string{"00:00", "06:00", "12:00", "18:00"},
			Production:  []float64{0, 0.9, 1.7, 0.4},
			Consumption: []float64{0.2, 5.0, 3.0, 0.8},
		},
		"chart-minus-4": {
			Labels:      []string{"00:00", "06:00", "12:00", "18:00"},
			Production:  []float64{0, 0.9, 1.7, 0.4},
			Consumption: []float64{0.2, 5.0, 3.0, 0.8},
		},
		"chart-minus-5": {
			Labels:      []string{"00:00", "06:00", "12:00", "18:00"},
			Production:  []float64{0, 0.9, 1.7, 0.4},
			Consumption: []float64{0.2, 5.0, 3.0, 0.8},
		},
		"chart-minus-6": {
			Labels:      []string{"00:00", "06:00", "12:00", "18:00"},
			Production:  []float64{0, 0.9, 1.7, 0.4},
			Consumption: []float64{0.2, 5.0, 3.0, 0.8},
		},
	}
	return hystoryData, nil
}
