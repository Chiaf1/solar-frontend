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
	data, err := s.repo.GetHistory()
	if err != nil {
		return nil, err
	}
	return data, nil
}
