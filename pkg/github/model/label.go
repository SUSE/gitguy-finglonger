package model

// LabelsPayload used to set new labels to a bug
type LabelsPayload struct {
	Labels []string `json:"labels"`
}
