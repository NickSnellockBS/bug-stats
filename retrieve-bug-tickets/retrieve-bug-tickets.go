package retrievebugtickets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
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

func RetrieveWorkflow(workflowId int) Workflow {
	var workflow Workflow

	client := &http.Client{}

	url := fmt.Sprintf("https://api.app.shortcut.com/api/v3/workflows/%d", workflowId)

	workflowReq, _ := http.NewRequest("GET", url, nil)
	workflowReq.Header.Add("Content-Type", "application/json")
	workflowReq.Header.Add("Shortcut-Token", "64f9d7a9-57ab-4376-a3fe-ad49d05ff641")

	workflowJson, _ := client.Do(workflowReq)
	workflowBody, _ := io.ReadAll(workflowJson.Body)

	json.Unmarshal(([]byte(workflowBody)), &workflow)

	return workflow
}

func RetrieveTickets(url string) string {
	client := &http.Client{}

	queryString := `type: bug`
	// queryString := `type: bug, created_at:1900-01-01..2023-07-01`
	// if completed {
	// 	queryString += ", completed:1900-01-01..2023-07-01"
	// }

	requestBody := strings.NewReader(fmt.Sprintf(`{"detail": "slim", "page_size": 1, "query": "%s"}`, queryString))
	fullUrl := fmt.Sprintf("https://api.app.shortcut.com%s", url)
	req, err := http.NewRequest("GET", fullUrl, requestBody)

	if (err != nil) {

	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Shortcut-Token", "64f9d7a9-57ab-4376-a3fe-ad49d05ff641")

	resp, err := client.Do(req)

	if (err != nil) {

	}

	body, err := io.ReadAll(resp.Body)

	if (err != nil) {

	}
	return string(body)
}