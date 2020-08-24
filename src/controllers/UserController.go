package controllers

import (
	"appstud.com/github-core/src/service"
	"github.com/gin-gonic/gin"
)

func UserController(engine *gin.Engine) {
	engine.GET("/api/users/register", handleRegisterUser)
	engine.GET("/api/users", handleGetAllUsers)
	engine.GET("/api/users/login", handleLoginUser)
	engine.GET("/api/users/me", handleConnectedUser)
}

func handleRegisterUser(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var password = c.Request.URL.Query().Get("password")
	var user, err = service.RegisterUser(username, password)
	if err == nil {
		c.JSON(201, user)
	} else {
		c.JSON(400, gin.H{"error message": err.Error()})
	}
}

func handleGetAllUsers(c *gin.Context) {
	c.JSON(200, service.GetAllUsers())
}

func handleLoginUser(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var password = c.Request.URL.Query().Get("password")
	var user, err = service.LoginUser(username, password)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error message": err.Error()})
	}
}

func handleConnectedUser(c *gin.Context) {
	var token = c.Request.URL.Query().Get("token")
	var user, err = service.GetConnectedUser(token)
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error message": err.Error()})
	}
}
