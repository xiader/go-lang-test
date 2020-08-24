package models

type UsernameTokenResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
