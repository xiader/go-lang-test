package models

type EasterEggResponse struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Timestamp int64  `json:"time"`
}
