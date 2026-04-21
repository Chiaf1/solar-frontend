package models

type ChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

type KPIData struct {
	ProductionValue  float64
	ProductionUnit   string
	ConsumptionValue float64
	ConsumptionUnit  string
}
