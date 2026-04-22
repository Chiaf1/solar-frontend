package services

import (
	"sync"
	"time"

	"github.com/chiaf1/solar-frontend/internal/models"
	"github.com/chiaf1/solar-frontend/internal/repositories"
)

type ChartService struct {
	repo repositories.EnergyRepository

	mu         sync.Mutex
	todayCache *cacheToday
	cacheTTL   time.Duration
}

type cacheToday struct {
	data      models.ChartData
	timestamp time.Time
}

func NewChartService(repo repositories.EnergyRepository) *ChartService {
	return &ChartService{
		repo:     repo,
		cacheTTL: 30 * time.Second,
	}
}

// Returns chart data for today's chart with cached data
func (s *ChartService) GetTodayChart() (models.ChartData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// use cache data if valid
	if s.todayCache != nil && time.Since(s.todayCache.timestamp) < s.cacheTTL {
		return s.todayCache.data, nil
	}

	// if not I call the API
	data, err := s.repo.GetToday()
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
	data, err := s.repo.GetHistory()
	if err != nil {
		return nil, err
	}
	return data, nil
}
