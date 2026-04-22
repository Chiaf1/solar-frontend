package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler for dashboard page
func (h *Handler) DashboardPage(ctx *gin.Context) {
	now := time.Now()

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

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		// Dati per grafici
		"ChartTodayJSON":     template.JS(todayJSON),
		"ChartYesterdayJSON": template.JS(historyChartsJSON["chart-yesterday"]),
		"ChartMinus2JSON":    template.JS(historyChartsJSON["chart-minus-2"]),
		"ChartMinus3JSON":    template.JS(historyChartsJSON["chart-minus-3"]),
		"ChartMinus4JSON":    template.JS(historyChartsJSON["chart-minus-4"]),
		"ChartMinus5JSON":    template.JS(historyChartsJSON["chart-minus-5"]),
		"ChartMinus6JSON":    template.JS(historyChartsJSON["chart-minus-6"]),

		// Data e ora
		"DayName":  now.Weekday().String(),
		"FullDate": now.Format("02 January 2006"),
		"Time":     now.Format("15:04"),

		// KPI
		"ProductionValue":  kpis.ProductionValue,
		"ProductionUnit":   kpis.ProductionUnit,
		"ConsumptionValue": kpis.ConsumptionValue,
		"ConsumptionUnit":  kpis.ConsumptionUnit,
	})
}
