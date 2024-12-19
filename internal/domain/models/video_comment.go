package models

type VideoComment struct {
	Timestamp

	ContentID     uint   `json:"contentID"`
	CourseGroupID string `json:"courseGroupID"`
	Comment       string `json:"comment"`
	UserID        uint   `json:"userID"`
}
