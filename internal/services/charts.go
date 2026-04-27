package services

import (
	"sync"
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
	"github.com/chiaf1/solar-frontend/internal/repositories"
)

type ChartService struct {
	energyRepo  repositories.EnergyRepository
	weatherRepo repositories.WeatherRepository

	todayMu       sync.Mutex
	todayCache    *cacheToday
	todayCacheTTL time.Duration

	weatherMu       sync.Mutex
	weatherCache    *cacheWeather
	weatherCacheTTL time.Duration
}

type cacheToday struct {
	data      models.ChartData
	timestamp time.Time
}

type cacheWeather struct {
	temperature float64
	timestamp   time.Time
}

func NewChartService(energyRepo repositories.EnergyRepository, weatherRepo repositories.WeatherRepository) *ChartService {
	return &ChartService{
		energyRepo:      energyRepo,
		weatherRepo:     weatherRepo,
		todayCacheTTL:   30 * time.Second,
		weatherCacheTTL: 10 * time.Minute,
	}
}

// Returns chart data for today's chart with cached data
func (s *ChartService) GetTodayChart() (models.ChartData, error) {
	s.todayMu.Lock()
	defer s.todayMu.Unlock()

	// use cache data if valid
	if s.todayCache != nil && time.Since(s.todayCache.timestamp) < s.todayCacheTTL {
		return s.todayCache.data, nil
	}

	// if not I call the API
	data, err := s.energyRepo.GetToday()
	if err != nil {
		// Fallback: if I have old cahce I use that
		if s.todayCache != nil {
			return s.todayCache.data, nil
		}
		return models.ChartData{}, err
	}

	// Update cache
	s.todayCache = &cacheToday{
		data:      data,
		timestamp: time.Now(),
	}

	return data, nil
}

// Returns chart data for history's charts
func (s *ChartService) GetHistoryCharts() (map[string]models.ChartData, error) {
	data, err := s.energyRepo.GetHistory()
	if err != nil {
		return nil, err
	}
	return data, nil
}
