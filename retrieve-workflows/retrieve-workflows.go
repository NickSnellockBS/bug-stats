package retrieveworkflows

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
)

type Workflow struct {
	Description string `json:"description"`
	EntityType  string `json:"entity_type"`
	ProjectIds  []int  `json:"project_ids"`
	States      []struct {
		Description       string    `json:"description"`
		EntityType        string    `json:"entity_type"`
		Verb              string    `json:"verb"`
		Name              string    `json:"name"`
		GlobalID          string    `json:"global_id"`
		NumStories        int       `json:"num_stories"`
		Type              string    `json:"type"`
		UpdatedAt         time.Time `json:"updated_at"`
		ID                int       `json:"id"`
		NumStoryTemplates int       `json:"num_story_templates"`
		Position          int       `json:"position"`
		CreatedAt         time.Time `json:"created_at"`
	} `json:"states"`
	Name            string    `json:"name"`
	UpdatedAt       time.Time `json:"updated_at"`
	AutoAssignOwner bool      `json:"auto_assign_owner"`
	ID              int       `json:"id"`
	TeamID          int       `json:"team_id"`
	CreatedAt       time.Time `json:"created_at"`
	DefaultStateID  int       `json:"default_state_id"`
}

func GetWorkflows() []Workflow {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	workflows := []Workflow{}

	client := &http.Client{}

	fullUrl := "https://api.app.shortcut.com/api/v3/workflows"

	req, err := http.NewRequest("GET", fullUrl, nil)

	if (err != nil) {
		fmt.Println(err)
		return workflows
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Shortcut-Token", os.Getenv("SHORTCUT_TOKEN"))

	resp, err := client.Do(req)

	if (err != nil) {
		fmt.Println(err)
		return workflows
	}

	body, err := io.ReadAll(resp.Body)

	if (err != nil) {
		fmt.Println(err)
		return workflows
	}

	json.Unmarshal(([]byte(body)), &workflows)

	return workflows
}