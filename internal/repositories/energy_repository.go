package repositories

import "github.com/chiaf1/solar-frontend/internal/models"

type EnergyRepository interface {
	GetToday() (models.ChartData, error)
	GetHistory() (map[string]models.ChartData, error)
	GetKPI() (models.ChartData, error)
}
