package services

import "github.com/chiaf1/solar-frontend/internal/models"

// Returns current KPIs values
func (s *ChartService) GetKPI() (models.KPIData, error) {
	return models.KPIData{
		ProductionValue:  2.3,
		ProductionUnit:   "KW",
		ConsumptionValue: 1.5,
		ConsumptionUnit:  "KW",
	}, nil
}
