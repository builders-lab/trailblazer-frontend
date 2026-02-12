package service

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func ForwardToBuilder(repo string) {
	payload := map[string]string{
		"repo_url": repo,
	}

	body, _ := json.Marshal(payload)

	http.Post(
		"http://localhost:8081/build",
		"application/json",
		bytes.NewBuffer(body),
	)
}
