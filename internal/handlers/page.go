package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for dashboard page
func (h *Handler) DashboardPage(ctx *gin.Context) {
	dateAndTime := h.service.GetDateAndTime()

	// Starting data
	todayChart, _ := h.service.GetTodayChart()
	historyCharts, _ := h.service.GetHistoryCharts()
	kpis, _ := h.service.GetKPI()

	// Json conversion
	todayJSON, _ := json.Marshal(todayChart)
	historyChartsJSON := make(map[string][]byte)
	for k := range historyCharts {
		historyChartsJSON[k], _ = json.Marshal(historyCharts[k])
	}
	temperature, _ := h.service.GetTemperature()

	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
		// Dati per grafici
		"ChartTodayJSON":     template.JS(todayJSON),
		"ChartYesterdayJSON": template.JS(historyChartsJSON["chart-yesterday"]),
		"ChartMinus2JSON":    template.JS(historyChartsJSON["chart-minus-2"]),
		"ChartMinus3JSON":    template.JS(historyChartsJSON["chart-minus-3"]),
		"ChartMinus4JSON":    template.JS(historyChartsJSON["chart-minus-4"]),
		"ChartMinus5JSON":    template.JS(historyChartsJSON["chart-minus-5"]),
		"ChartMinus6JSON":    template.JS(historyChartsJSON["chart-minus-6"]),

		// Data e ora
		"DayName":     dateAndTime.DayName,
		"FullDate":    dateAndTime.Date,
		"Time":        dateAndTime.Time,
		"Temperature": temperature,

		// KPI
		"ProductionValue":  kpis.ProductionValue,
		"ProductionUnit":   kpis.ProductionUnit,
		"ConsumptionValue": kpis.ConsumptionValue,
		"ConsumptionUnit":  kpis.ConsumptionUnit,
	})
}

// Handler for today page
func (h *Handler) TodayPage(ctx *gin.Context) {
	dateAndTime := h.service.GetDateAndTime()

	// Starting data
	todayChart, _ := h.service.GetTodayChart()
	kpis, _ := h.service.GetKPI()
	temperature, _ := h.service.GetTemperature()

	// Json conversion
	todayJSON, _ := json.Marshal(todayChart)

	ctx.HTML(http.StatusOK, "today_page.html", gin.H{
		// Data e ora
		"DayName":     dateAndTime.DayName,
		"FullDate":    dateAndTime.Date,
		"Time":        dateAndTime.Time,
		"Temperature": temperature,

		// KPI
		"ProductionValue":  kpis.ProductionValue,
		"ProductionUnit":   kpis.ProductionUnit,
		"ConsumptionValue": kpis.ConsumptionValue,
		"ConsumptionUnit":  kpis.ConsumptionUnit,

		// Dati per grafico
		"ChartTodayJSON": template.JS(todayJSON),
	})
}
