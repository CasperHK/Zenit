package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// registerAPIRoutes registers API routes under /api
func registerAPIRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

