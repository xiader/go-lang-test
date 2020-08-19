package controllers

import (
	"appstud.com/github-core/src/service"
	"github.com/gin-gonic/gin"
)

func HealthCheckController(engine *gin.Engine) {
	engine.GET("/api/healthcheck", handleHealth)
}

func handleHealth(c *gin.Context) {
	c.JSON(200, service.GetHealthCheck())
}
