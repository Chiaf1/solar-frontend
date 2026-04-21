package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chiaf1/solar-frontend/internal/services"
	"github.com/gin-gonic/gin"
)

// Handler for today chart data refresh
func RefreshToday(ctx *gin.Context) {
	data := services.GetTodayChart()

	// Prepara il payload per HTMX
	payload := map[string]any{
		"updateChartToday": data,
	}
	payloadJSON, _ := json.Marshal(payload)

	// Invia i dati tramite header (HTMX li intercetterà)
	ctx.Header("HX-Trigger", string(payloadJSON))
	ctx.Status(http.StatusNoContent)
}

// Handler for history charts data refresh
func RefreshHistory(ctx *gin.Context) {
	data := services.GetHistoryCharts()
	// Prepara il payload per HTMX
	payload := map[string]any{
		"updateChartHistory": data,
	}
	payloadJSON, _ := json.Marshal(payload)

	// Invia i dati tramite header (HTMX li intercetterà)
	ctx.Header("HX-Trigger", string(payloadJSON))
	ctx.Status(http.StatusNoContent)
}
