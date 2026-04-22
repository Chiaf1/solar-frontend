package models

type ChartData struct {
	Labels      []string   `json:"labels"`
	Production  []*float64 `json:"production"`
	Consumption []*float64 `json:"consumption"`
}

type ChartPoint struct {
	Timestamp   string
	Production  float64
	Consumption float64
}

type KPIData struct {
	ProductionValue  float64
	ProductionUnit   string
	ConsumptionValue float64
	ConsumptionUnit  string
}

type DateAndTime struct {
	DayName string
	Date    string
	Time    string
}

var Giorni = []string{
	"Domenica",
	"Lunedì",
	"Martedì",
	"Mercoledì",
	"Giovedì",
	"Sabato",
}

var Mesi = []string{
	"Gennaio",
	"Febbraio",
	"Marzo",
	"Aprile",
	"Maggio",
	"Giugno",
	"Luglio",
	"Agosto",
	"Settembre",
	"Ottobre",
	"Novembre",
	"Dicembre",
}
