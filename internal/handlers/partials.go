package handlers

import (
	"net/http"
	"time"

	"github.com/chiaf1/solar-frontend/internal/services"
	"github.com/gin-gonic/gin"
)

// Handler for header values
func HeaderPartial(ctx *gin.Context) {
	now := time.Now()

	ctx.HTML(http.StatusOK, "partials/header", gin.H{
		"DayName":  now.Weekday().String(),
		"FullDate": now.Format("02 January 2006"),
		"Time":     now.Format("15:04"),
	})
}

// Handler for KPIs values
func KPIsPartial(ctx *gin.Context) {
	data := services.GetKPI()

	ctx.HTML(http.StatusOK, "partials/kpis", gin.H{
		// Production
		"ProductionValue": data.ProductionValue,
		"ProductionUnit":  data.ProductionUnit,
		// Consumption
		"ConsumptionValue": data.ConsumptionValue,
		"ConsumptionUnit":  data.ConsumptionUnit,
	})
}
