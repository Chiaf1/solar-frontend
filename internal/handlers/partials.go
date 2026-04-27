package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler for header values
func (h *Handler) HeaderPartial(ctx *gin.Context) {
	dateAndTime := h.service.GetDateAndTime()
	temp, _ := h.service.GetTemperature()

	ctx.HTML(http.StatusOK, "partials/header", gin.H{
		"DayName":     dateAndTime.DayName,
		"FullDate":    dateAndTime.Date,
		"Time":        dateAndTime.Time,
		"Temperature": temp,
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

// Handler for temperature value
func (h *Handler) TemperaturePartial(ctx *gin.Context) {
	data, err := h.service.GetTemperature()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch temperature data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"temperatura": data,
	})
}
