package models

type Repository struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
	HTMLURL  string `json:"html_url"`
}

type Author struct {
	Name string `json:"name"`
}

type Commit struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Author  Author `json:"author"`
}

type PushEvent struct {
	Ref        string     `json:"ref"`
	Repository Repository `json:"repository"`
	Pusher     Author     `json:"pusher"`
	Commits    []Commit   `json:"commits"`
}
