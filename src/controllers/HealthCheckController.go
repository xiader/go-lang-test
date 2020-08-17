package controllers

import (
	"appstud.com/github-core/src/models"
	"github.com/gin-gonic/gin"
)

// HealthCheckController - Route controller
func HealthCheckController(engine *gin.Engine) {
	engine.GET("/api/healthcheck", handleHealth)
}

func handleHealth(c *gin.Context) {
	c.JSON(200, models.HealthCheckResponse{
		Name:    "github-api",
		Version: "1.0",
	})
}
