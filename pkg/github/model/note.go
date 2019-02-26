package model

import "time"

// NotePayload payload to send to the project
type NotePayload struct {
	ContentID   int    `json:"content_id"`
	ContentType string `json:"content_type"`
	Note        string `json:"note,omitempty"`
}

// MoveNotePayload payload to move to In Progress column
type MoveNotePayload struct {
	Position string `json:"position"`
	ColumnID int    `json:"column_id"`
}

// NoteInColumns return struct for all the notes in a project column
type NoteInColumns struct {
	URL        string      `json:"url"`
	ProjectURL string      `json:"project_url"`
	ID         int         `json:"id"`
	NodeID     string      `json:"node_id"`
	Note       interface{} `json:"note"`
	Archived   bool        `json:"archived"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	ColumnURL  string      `json:"column_url"`
	ContentURL string      `json:"content_url,omitempty"`
}
