package models

type TermsAndCondition struct {
	Timestamp
	Title       string `json:"title"`
	Description string `json:"description"`
}
