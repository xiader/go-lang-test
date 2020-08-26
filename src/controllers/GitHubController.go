package controllers

import (
	"appstud.com/github-core/src/service"
	"github.com/gin-gonic/gin"
)

func GitHubController(engine *gin.Engine) {
	engine.GET("/api/github/feed", handleGitHubFeed)
}

func handleGitHubFeed(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var feed, err = service.GetGitHubFeedForUser(username)
	if err != nil {
		c.JSON(400, gin.H{"error message": err.Error()})
	} else {
		c.JSON(200, feed)
	}
}
