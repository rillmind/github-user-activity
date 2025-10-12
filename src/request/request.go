package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubEvent struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action   string   `json:"action"`
	Ref_type string   `json:"ref_type"`
	Commits  []Commit `json:"commits"`
}

type Commit struct {
	Sha string `json:"sha"`
}

func Request(username string) {
	var events []GitHubEvent

	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&events)

	if err != nil {
		fmt.Println(err)
	}

	for _, event := range events {
		switch event.Type {
		case "CreateEvent":
			fmt.Println("-", "Created repository", event.Repo.Name)
		case "PushEvent":
			if len(event.Payload.Commits) == 1 {
				fmt.Println("-", "Pushed", len(event.Payload.Commits), "commit to", event.Repo.Name)
			} else {
				fmt.Println("-", "Pushed", len(event.Payload.Commits), "commits to", event.Repo.Name)
			}
		case "WatchEvent":
			if event.Payload.Action == "started" {
				fmt.Println("-", "Started", event.Repo.Name)
			}
		}
	}
}
