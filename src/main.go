package main

import (
	"appstud.com/github-core/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.HelloWorldController(r)
	controllers.HealthCheckController(r)
	controllers.EasterEggController(r)
	r.Run()
}
