package services

import "github.com/chiaf1/solar-frontend/internal/models"

// Returns current KPIs values
func (s *ChartService) GetKPI() (models.KPIData, error) {
	today, err := s.GetTodayChart()
	if err != nil {
		return models.KPIData{}, err
	}
	prod := lastNotNil(today.Production)
	cons := lastNotNil(today.Consumption)

	return models.KPIData{
		ProductionValue:  *prod,
		ProductionUnit:   "KW",
		ConsumptionValue: *cons,
		ConsumptionUnit:  "KW",
	}, nil
}

// Returns last not nil value from a pointer to a slice of float64
func lastNotNil(vals []*float64) *float64 {
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] != nil {
			return vals[i]
		}
	}
	return nil
}
