package service

import (
	"appstud.com/github-core/src/models"
	"time"
)

func GetHealthCheck() models.HealthCheckResponse {
	return eggs
}

var eggs = models.HealthCheckResponse{
	Name:      "github-api",
	Version:   "1.0",
	Timestamp: time.Now().Unix(),
}
