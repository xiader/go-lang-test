package models

type HealthCheckResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
