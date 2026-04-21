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

	//r.LoadHTMLGlob("web/templates/*")
	r.LoadHTMLFiles(
		"web/templates/base.html",
		"web/templates/index.html",
		"web/templates/partials/chart_today.html",
		"web/templates/partials/chart_yesterday.html",
		"web/templates/partials/header.html",
		"web/templates/partials/kpi_production.html",
		"web/templates/partials/kpi_consumption.html",
		"web/templates/partials/kpis.html",
		"web/templates/partials/chart_minus_2.html",
		"web/templates/partials/chart_minus_3.html",
		"web/templates/partials/chart_minus_4.html",
		"web/templates/partials/chart_minus_5.html",
		"web/templates/partials/chart_minus_6.html",
	)
	r.Static("/static", "./web/static")

	// Endpoint pagina principale
	r.GET("/", func(ctx *gin.Context) {
		now := time.Now()
		ctx.HTML(200, "index.html", gin.H{
			// Dati per grafici
			"ChartTodayJSON":     template.JS(todayJSON),
			"ChartYesterdayJSON": template.JS(yesterdayJSON),
			"ChartMinus2JSON":    template.JS(yesterdayJSON),
			"ChartMinus3JSON":    template.JS(yesterdayJSON),
			"ChartMinus4JSON":    template.JS(yesterdayJSON),
			"ChartMinus5JSON":    template.JS(yesterdayJSON),
			"ChartMinus6JSON":    template.JS(yesterdayJSON),
			"ChartMinus7JSON":    template.JS(yesterdayJSON),

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

	// Endpoint per refresh grafico oggi
	r.GET("/api/refresh-today", func(ctx *gin.Context) {
		// Recupera i dati aggiornati
		newData := ChartData{
			Labels: []string{"00:00", "04:00", "08:00", "12:00", "16:00", "20:00"},
			Values: []float64{0.1, 0.2, 1.5, 4.2, 2.1, 0.5},
		}

		// Prepara il payload per HTMX
		payload := map[string]interface{}{
			"updateChartToday": newData,
		}
		payloadJSON, _ := json.Marshal(payload)

		// Invia i dati tramite header (HTMX li intercetterà)
		ctx.Header("HX-Trigger", string(payloadJSON))

		ctx.Status(204)
	})

	// Endpoint per regresh grafico ieri
	r.GET("/api/refresh-yesterday", func(ctx *gin.Context) {
		// Recupera i dati aggiornati
		newData := ChartData{
			Labels: []string{"00:00", "04:00", "08:00", "12:00", "16:00", "20:00"},
			Values: []float64{0.1, 0.2, 1.5, 4.2, 2.1, 0.5},
		}

		// Prepara il payload per HTMX
		payload := map[string]interface{}{
			"updateChartYesterday": newData,
		}
		payloadJSON, _ := json.Marshal(payload)

		// Invia i dati tramite header (HTMX li intercetterà)
		ctx.Header("HX-Trigger", string(payloadJSON))

		ctx.Status(204)
	})

	// Endpoint per refrash grafici history ovvero da ieri fino a -6
	r.GET("/api/refresh-history", func(ctx *gin.Context) {
		// Recupera i dati aggiornati
		hystoryData := map[string]ChartData{
			"chart-yesterday": {
				Labels: []string{"00:00", "04:00", "08:00", "12:00", "16:00", "20:00"},
				Values: []float64{0.5, 1.2, 2.3, 1.8, 0.9, 0.3},
			},
			"chart-minus-2": {
				Labels: []string{"00:00", "06:00", "12:00", "18:00"},
				Values: []float64{0.2, 1.1, 2.0, 0.8},
			},
			"chart-minus-3": {
				Labels: []string{"00:00", "06:00", "12:00", "18:00"},
				Values: []float64{0, 0.9, 1.7, 0.4},
			},
			"chart-minus-4": {
				Labels: []string{"00:00", "06:00", "12:00", "18:00"},
				Values: []float64{0, 0.9, 1.7, 0.4},
			},
			"chart-minus-5": {
				Labels: []string{"00:00", "06:00", "12:00", "18:00"},
				Values: []float64{0, 0.9, 1.7, 0.4},
			},
			"chart-minus-6": {
				Labels: []string{"00:00", "06:00", "12:00", "18:00"},
				Values: []float64{0, 0.9, 1.7, 0.4},
			},
		}

		// Prepara il payload per HTMX
		payload := map[string]any{
			"updateChartHistory": hystoryData,
		}
		payloadJSON, _ := json.Marshal(payload)

		// Invia i dati tramite header (HTMX li intercetterà)
		ctx.Header("HX-Trigger", string(payloadJSON))

		ctx.Status(204)
	})

	// Endpoint for date and time update
	r.GET("/partials/header", func(ctx *gin.Context) {
		now := time.Now()

		ctx.HTML(200, "partials/header", gin.H{
			"DayName":  now.Weekday().String(),
			"FullDate": now.Format("02 January 2006"),
			"Time":     now.Format("15:04"),
		})
	})

	// Endpoint for KPIs update
	r.GET("/partials/kpis", func(ctx *gin.Context) {
		currentProduction += 1
		ctx.HTML(200, "partials/kpis", gin.H{
			// Production
			"ProductionValue": currentProduction,
			"ProductionUnit":  "kW",
			// Consumption
			"ConsumptionValue": 1.50,
			"ConsumptionUnit":  "kW",
		})
	})

	r.Run(":8080")
}
