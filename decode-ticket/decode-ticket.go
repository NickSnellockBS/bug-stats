package decodeticket

import (
	"encoding/json"
	"time"
)

type Ticket struct {
	Next string `json:"next"`
	Data []struct {
		AppURL     string `json:"app_url"`
		Archived   bool   `json:"archived"`
		Started    bool   `json:"started"`
		StoryLinks []struct {
			EntityType             string    `json:"entity_type"`
			ObjectID               int       `json:"object_id"`
			Verb                   string    `json:"verb"`
			Type                   string    `json:"type"`
			UpdatedAt              time.Time `json:"updated_at"`
			ID                     int       `json:"id"`
			SubjectID              int       `json:"subject_id"`
			SubjectWorkflowStateID int       `json:"subject_workflow_state_id"`
			CreatedAt              time.Time `json:"created_at"`
		} `json:"story_links"`
		EntityType string `json:"entity_type"`
		Labels     []struct {
			AppURL      string      `json:"app_url"`
			Description interface{} `json:"description"`
			Archived    bool        `json:"archived"`
			EntityType  string      `json:"entity_type"`
			Color       interface{} `json:"color"`
			Name        string      `json:"name"`
			GlobalID    string      `json:"global_id"`
			UpdatedAt   time.Time   `json:"updated_at"`
			ExternalID  interface{} `json:"external_id"`
			ID          int         `json:"id"`
			CreatedAt   time.Time   `json:"created_at"`
		} `json:"labels"`
		TaskIds              []interface{} `json:"task_ids"`
		MentionIds           []interface{} `json:"mention_ids"`
		MemberMentionIds     []interface{} `json:"member_mention_ids"`
		StoryType            string        `json:"story_type"`
		CustomFields         []interface{} `json:"custom_fields"`
		FileIds              []interface{} `json:"file_ids"`
		NumTasksCompleted    int           `json:"num_tasks_completed"`
		WorkflowID           int           `json:"workflow_id"`
		CompletedAtOverride  interface{}   `json:"completed_at_override"`
		StartedAt            interface{}   `json:"started_at"`
		CompletedAt          interface{}   `json:"completed_at"`
		Name                 string        `json:"name"`
		GlobalID             string        `json:"global_id"`
		Completed            bool          `json:"completed"`
		Blocker              bool          `json:"blocker"`
		EpicID               int           `json:"epic_id"`
		StoryTemplateID      interface{}   `json:"story_template_id"`
		ExternalLinks        []interface{} `json:"external_links"`
		PreviousIterationIds []interface{} `json:"previous_iteration_ids"`
		RequestedByID        string        `json:"requested_by_id"`
		IterationID          interface{}   `json:"iteration_id"`
		LabelIds             []int         `json:"label_ids"`
		StartedAtOverride    interface{}   `json:"started_at_override"`
		GroupID              string        `json:"group_id"`
		WorkflowStateID      int           `json:"workflow_state_id"`
		UpdatedAt            time.Time     `json:"updated_at"`
		GroupMentionIds      []interface{} `json:"group_mention_ids"`
		FollowerIds          []string      `json:"follower_ids"`
		OwnerIds             []interface{} `json:"owner_ids"`
		ExternalID           interface{}   `json:"external_id"`
		ID                   int           `json:"id"`
		Estimate             interface{}   `json:"estimate"`
		Position             int64         `json:"position"`
		Blocked              bool          `json:"blocked"`
		ProjectID            int           `json:"project_id"`
		LinkedFileIds        []interface{} `json:"linked_file_ids"`
		Deadline             interface{}   `json:"deadline"`
		Stats                struct {
			NumRelatedDocuments int `json:"num_related_documents"`
		} `json:"stats"`
		CommentIds []interface{} `json:"comment_ids"`
		CreatedAt  time.Time     `json:"created_at"`
		MovedAt    time.Time     `json:"moved_at"`
		LeadTime   int           `json:"lead_time,omitempty"`
		CycleTime  int           `json:"cycle_time,omitempty"`
	} `json:"data"`
	Total int `json:"total"`
}

func DecodeTicket(ticketResponse string) Ticket {
	var ticket Ticket

	json.Unmarshal(([]byte(ticketResponse)), &ticket)
	return ticket
}