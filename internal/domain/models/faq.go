package models

type FAQ struct {
	Timestamp
	Title       string `json:"title"`
	Description string `json:"description"`
}
