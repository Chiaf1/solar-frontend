package main

import (
	"encoding/json"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

type ChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

func main() {
	// Dati finti per grafici
	// Grafico oggi
	todayData := ChartData{
		Labels: []string{
			"00:00", "02:00", "04:00", "06:00",
			"08:00", "10:00", "12:00", "14:00",
			"16:00", "18:00", "20:00", "22:00",
		},
		Values: []float64{
			0, 0, 0, 0.5,
			1.2, 2.8, 4.1, 3.9,
			2.6, 1.1, 0.3, 1,
		},
	}
	todayJSON, _ := json.Marshal(todayData)

	//Grafico ieri
	yesterdayData := ChartData{
		Labels: []string{
			"00:00", "02:00", "04:00", "06:00",
			"08:00", "10:00", "12:00", "14:00",
			"16:00", "18:00", "20:00", "22:00",
		},
		Values: []float64{
			0, 0, 0, 0.3,
			1.0, 2.4, 3.6, 3.2,
			2.0, 0.9, 0.2, 0,
		},
	}
	yesterdayJSON, _ := json.Marshal(yesterdayData)

	// Estrapolazione KPI da dati di oggi
	currentProduction := 0.0
	if len(todayData.Values) > 0 {
		currentProduction = todayData.Values[len(todayData.Values)-1]
	}

	// Creazione router gin
	r := gin.Default()

	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "./web/static")

	r.GET("/", func(ctx *gin.Context) {
		now := time.Now()
		ctx.HTML(200, "index.html", gin.H{
			// Dati per grafici
			"ChartTodayJSON":     template.JS(todayJSON),
			"ChartYesterdayJSON": template.JS(yesterdayJSON),

			// Data e ora
			"DayName":  now.Weekday().String(),
			"FullDate": now.Format("02 January 2006"),
			"Time":     now.Format("15:04"),

			// KPI finti
			"ProductionValue":  currentProduction,
			"ProductionUnit":   "kW",
			"ConsumptionValue": 1.87,
			"ConsumptionUnit":  "kW",
		})
	})

	r.Run(":8080")
}
