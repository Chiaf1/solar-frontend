package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for header values
func (h *Handler) HeaderPartial(ctx *gin.Context) {
	dateAndTime := h.service.GetDateAndTime()

	ctx.HTML(http.StatusOK, "partials/header", gin.H{
		"DayName":  dateAndTime.DayName,
		"FullDate": dateAndTime.Date,
		"Time":     dateAndTime.Time,
	})
}

// Handler for KPIs values
func (h *Handler) KPIsPartial(ctx *gin.Context) {
	data, _ := h.service.GetKPI()

	ctx.HTML(http.StatusOK, "partials/kpis", gin.H{
		// Production
		"ProductionValue": data.ProductionValue,
		"ProductionUnit":  data.ProductionUnit,
		// Consumption
		"ConsumptionValue": data.ConsumptionValue,
		"ConsumptionUnit":  data.ConsumptionUnit,
	})
}
