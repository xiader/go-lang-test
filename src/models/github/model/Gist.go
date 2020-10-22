package model

import "time"

type Gist struct {
	URL         string      `json:"url"`
	ForksURL    string      `json:"forks_url"`
	CommitsURL  string      `json:"commits_url"`
	ID          string      `json:"id"`
	NodeID      string      `json:"node_id"`
	GitPullURL  string      `json:"git_pull_url"`
	GitPushURL  string      `json:"git_push_url"`
	HTMLURL     string      `json:"html_url"`
	Files       interface{} `json:"files"`
	Public      bool        `json:"public"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Description string      `json:"description"`
	Comments    int         `json:"comments"`
	User        interface{} `json:"user"`
	CommentsURL string      `json:"comments_url"`
	Owner       `json:"owner"`
	Truncated   bool `json:"truncated"`
}
