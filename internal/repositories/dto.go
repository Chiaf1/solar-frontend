package repositories

type EnergyAPIPoint struct {
	Time        string  `json:"time"`
	Production  float64 `json:"production"`
	Consumption float64 `json:"consumption"`
}

type EnergyAPIDaily struct {
	Day    string           `json:"day"`
	Points []EnergyAPIPoint `json:"points"`
}
