package repositories

import "time"

type EnergyAPIPoint struct {
	Time        time.Time `json:"time"`
	Production  float64   `json:"production"`
	Consumption float64   `json:"consumption"`
}

type EnergyAPIDaily struct {
	Day    string           `json:"day"`
	Points []EnergyAPIPoint `json:"points"`
}
