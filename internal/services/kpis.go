package services

import "github.com/chiaf1/solar-frontend/internal/models"

// Returns current KPIs values
func GetKPI() models.KPIData {
	return models.KPIData{
		ProductionValue:  2.3,
		ProductionUnit:   "KW",
		ConsumptionValue: 1.5,
		ConsumptionUnit:  "KW",
	}
}
