package handlers

import "github.com/gin-gonic/gin"

// Register all page routes
func (h *Handler) RegisterPageRoutes(r *gin.Engine) {
	r.GET("/", h.DashboardPage)
	r.GET("/dashboard", h.DashboardPage)
}

// Register all API routes
func (h *Handler) RegisterApiRoutes(r *gin.Engine) {
	r.GET("/api/refresh-today", h.RefreshToday)
	r.GET("/api/refresh-history", h.RefreshHistory)
}

// Register all partials routes
func (h *Handler) RegisterPartialRoutes(r *gin.Engine) {
	r.GET("/partials/header", h.HeaderPartial)
	r.GET("/partials/kpis", h.KPIsPartial)
}
