package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chiaf1/solar-frontend/internal/handlers"
	"github.com/chiaf1/solar-frontend/internal/repositories"
	"github.com/chiaf1/solar-frontend/internal/services"
	"github.com/gin-gonic/gin"
)

const defaultURL = "http://192.168.0.171:8080"

func main() {
	// Get base url from env
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = defaultURL
	}

	// Creating the repository, service and handler layers to retrive the API data
	repo := repositories.NewEnergyAPIRepository(baseURL)
	service := services.NewChartService(repo)
	handler := handlers.NewHandler(service)

	// Creazione router gin
	r := gin.Default()

	//r.LoadHTMLGlob("web/templates/*")
	r.LoadHTMLFiles(
		"web/templates/base.html",
		"web/templates/base_partials.html",
		"web/templates/pages/dashboard.html",
		"web/templates/pages/today_page.html",
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

	// Creating an HTTP server to handle graceful shutdowns
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Starting the server in a go routine
	go func() {
		log.Println("Server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()

	//Waiting for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited cleanly")
}
