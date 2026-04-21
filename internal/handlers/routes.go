package handlers

import "github.com/gin-gonic/gin"

// Register all page routes
func RegisterPageRoutes(r *gin.Engine) {
	r.GET("/", DashboardPage)
}

// Register all API routes
func RegisterApiRoutes(r *gin.Engine) {
	r.GET("/api/refresh-today", RefreshToday)
	r.GET("/api/refresh-history", RefreshHistory)
}

// Register all partials routes
func RegisterPartialRoutes(r *gin.Engine) {
	r.GET("/partials/header", HeaderPartial)
	r.GET("/partials/kpis", KPIsPartial)
}
