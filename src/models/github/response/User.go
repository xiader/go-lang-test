package response

type User struct {
	ID      int64   `json:"id"`
	Login   string  `json:"login"`
	Avatar  string  `json:"avatar"`
	Details Details `json:"details"`
}

type Details struct {
	PublicRepos []PublicRepo    `json:"public_repos"`
	PublicGists []PublicGist    `json:"public_gists"`
	Followers   []BriefUserInfo `json:"followers"`
	Following   []BriefUserInfo `json:"following"`
}

type PublicRepo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PublicGist struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

type BriefUserInfo struct {
	Login  string `json:"login"`
	ID     int64  `json:"id"`
	Avatar string `json:"avatar_id"`
	Url    string `json:"url"`
}
