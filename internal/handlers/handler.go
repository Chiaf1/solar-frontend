package handlers

import "github.com/chiaf1/solar-frontend/internal/services"

type Handler struct {
	service *services.ChartService
}

// Return new handler struct
func NewHandler(service *services.ChartService) *Handler {
	return &Handler{
		service: service,
	}
}
