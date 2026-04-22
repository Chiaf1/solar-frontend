package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for today chart data refresh
func (h *Handler) RefreshToday(ctx *gin.Context) {
	data, _ := h.service.GetTodayChart()

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
func (h *Handler) RefreshHistory(ctx *gin.Context) {
	data, _ := h.service.GetHistoryCharts()
	// Prepara il payload per HTMX
	payload := map[string]any{
		"updateChartHistory": data,
	}
	payloadJSON, _ := json.Marshal(payload)

	// Invia i dati tramite header (HTMX li intercetterà)
	ctx.Header("HX-Trigger", string(payloadJSON))
	ctx.Status(http.StatusNoContent)
}
