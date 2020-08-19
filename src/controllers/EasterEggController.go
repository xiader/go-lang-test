package controllers

import (
	"appstud.com/github-core/src/service"
	"github.com/gin-gonic/gin"
)

func EasterEggController(engine *gin.Engine) {
	engine.GET("/api/timemachine/logs/mcfly", handleEasterEgg)
}

func handleEasterEgg(c *gin.Context) {
	c.JSON(200, service.GetEasterEggs())
}
