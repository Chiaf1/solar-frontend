package main

import (
	"github.com/chiaf1/solar-frontend/internal/handlers"
	"github.com/gin-gonic/gin"
)

type ChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
}

func main() {
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

	handlers.RegisterPageRoutes(r)
	handlers.RegisterPartialRoutes(r)
	handlers.RegisterApiRoutes(r)

	r.Run(":8080")
}
