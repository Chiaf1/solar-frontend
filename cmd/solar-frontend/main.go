package main

import (
	"github.com/chiaf1/solar-frontend/internal/handlers"
	"github.com/chiaf1/solar-frontend/internal/repositories"
	"github.com/chiaf1/solar-frontend/internal/services"
	"github.com/gin-gonic/gin"
)

const baseURL = "http://192.168.0.171:8080"

func main() {

	// Creating the repository, service and handler layers to retrive the API data
	repo := repositories.NewEnergyAPIRepository(baseURL)
	service := services.NewChartService(repo)
	handler := handlers.NewHandler(service)

	// Creazione router gin
	r := gin.Default()

	//r.LoadHTMLGlob("web/templates/*")
	r.LoadHTMLFiles(
		"web/templates/base.html",
		"web/templates/pages/dashboard.html",
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

	handler.RegisterPageRoutes(r)
	handler.RegisterPartialRoutes(r)
	handler.RegisterApiRoutes(r)

	r.Run(":8080")
}
