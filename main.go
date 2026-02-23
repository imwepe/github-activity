package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Event struct {
	Type      string          `json:"type"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt time.Time       `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type PushPayload struct {
	Commits []struct {
		Message string `json:"message"`
	} `json:"commits"`
}

type IssuesPayload struct {
	Action string `json:"action"`
}

type WatchPayload struct {
	Action string `json:"action"`
}

func main() {
	var username string

	fmt.Print("Masukkan username github: ")
	fmt.Scanln(&username)

	if username == "" {
		fmt.Println("Username can't be empty")
		return
	}

	events, err := fetchGitHubEvents(username)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for {
		var filter string
		var choice int

		fmt.Println("\nChoose Github activity:")
		fmt.Println("1. PushEvent")
		fmt.Println("2. PullRequestEvent")
		fmt.Println("3. IssuesEvent")
		fmt.Println("4. WatchEvent")
		fmt.Println("5. All")
		fmt.Println("6. Exit")
		fmt.Print("Input selection: ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			filter = "PushEvent"
		case 2:
			filter = "PullRequestEvent"
		case 3:
			filter = "IssuesEvent"
		case 4:
			filter = "WatchEvent"
		case 5:
			filter = "all"
		case 6:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid.")
			continue
		}

		filtered := filterEvents(events, filter)

		if len(filtered) == 0 {
			fmt.Println("No events found.")
			continue
		}

		fmt.Println("==== Latest Github Activities ====")
		limit := 5
		if len(filtered) < 5 {
			limit = len(filtered)
		}

		for i := 0; i < limit; i++ {
			e := filtered[i]

			switch e.Type {

			case "PushEvent":
				var payload PushPayload
				json.Unmarshal(e.Payload, &payload)
				fmt.Printf("- Pushed %d commits to %s\n",
					len(payload.Commits),
					e.Repo.Name)

			case "IssuesEvent":
				var payload IssuesPayload
				json.Unmarshal(e.Payload, &payload)
				if payload.Action == "opened" {
					fmt.Printf("- Opened a new issue in %s\n", e.Repo.Name)
				}

			case "WatchEvent":
				fmt.Printf("- Starred %s\n", e.Repo.Name)

			default:
				fmt.Printf("- %s in %s\n", e.Type, e.Repo.Name)
			}
		}
	}
}

// fetching from github
func fetchGitHubEvents(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// make http get and error handling for http request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Github API")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("user not found")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Github API error: %s", resp.Status)
	}

	var events []Event
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response")
	}

	return events, nil
}

// filtering event types
func filterEvents(events []Event, filter string) []Event {
	if strings.ToLower(filter) == "all" {
		return events
	}

	var result []Event
	for _, e := range events {
		if strings.EqualFold(e.Type, filter) {
			result = append(result, e)
		}
	}
	return result
}
