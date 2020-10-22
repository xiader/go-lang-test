package response

type Event struct {
	Type  string    `json:"type"`
	Actor Actor     `json:"actor"`
	Repo  EventRepo `json:"repo"`
}

type Actor struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type EventRepo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
