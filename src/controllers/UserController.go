package controllers

import (
	"appstud.com/github-core/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserController(engine *gin.Engine) {
	engine.GET("/api/users/register", handleRegisterUser)
}

func handleRegisterUser(c *gin.Context) {
	var username = c.Request.URL.Query().Get("username")
	var password = c.Request.URL.Query().Get("password")
	fmt.Println(username + "__ " + password)
	var user, err = service.RegisterUser(username, password)
	if err == nil {
		c.JSON(201, user)
	} else {
		c.JSON(400, gin.H{"error message": err.Error()})
	}
}
