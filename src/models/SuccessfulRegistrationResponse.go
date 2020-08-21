package models

type SuccessfulRegistrationResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
