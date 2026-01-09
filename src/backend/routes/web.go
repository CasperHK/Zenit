package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	middleware "backend/app/Http/Middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all web and API routes and middlewares.
func RegisterRoutes(router *gin.Engine) {
	router.Use(middleware.RequestCounter())

	// Web routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// metrics
	router.GET("/metrics", func(c *gin.Context) {
		c.Header("Content-Type", "text/plain; version=0.0.4")
		count := middleware.GetRequestCount()
		fmt.Fprintf(c.Writer, "# HELP zenit_requests_total Total HTTP requests processed by the server\n")
		fmt.Fprintf(c.Writer, "# TYPE zenit_requests_total counter\n")
		fmt.Fprintf(c.Writer, "zenit_requests_total %d\n", count)
		fmt.Fprintf(c.Writer, "# HELP zenit_build_info Build info\n")
		fmt.Fprintf(c.Writer, "# TYPE zenit_build_info gauge\n")
		fmt.Fprintf(c.Writer, "zenit_build_info{version=\"dev\"} 1\n")
	})

	// Serve static files from ./public if available
	publicDir := filepath.Join("public")
	if stat, err := os.Stat(publicDir); err == nil && stat.IsDir() {
		router.StaticFS("/", http.Dir(publicDir))
	}

	// API routes (defined in api.go)
	registerAPIRoutes(router)
}
