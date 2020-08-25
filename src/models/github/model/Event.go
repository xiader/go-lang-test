package model

import "time"

type Event struct {
	Id        string                 `json:"id"`
	Type      string                 `json:"type"`
	Actor     Actor                  `json:"actor"`
	Repo      Repo                   `json:"repo"`
	Payload   map[string]interface{} `json:"payload"`
	Public    bool                   `json:"public"`
	CreatedAt time.Time              `json:"created_at"`
}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
