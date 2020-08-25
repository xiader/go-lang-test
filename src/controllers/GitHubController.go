package controllers

import (
	"appstud.com/github-core/src/service"
	"github.com/gin-gonic/gin"
)

func GitHubController(engine *gin.Engine) {
	engine.GET("/api/github/feed", handleGitHubFeed)
}

func handleGitHubFeed(c *gin.Context) {
	var feed, err = service.GetGitHubFeed()
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, feed)
	}
}
