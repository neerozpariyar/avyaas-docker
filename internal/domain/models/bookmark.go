package models

type Bookmark struct {
	Timestamp
	ContentID    uint   `json:"contentID"`
	QuestionID   uint   `json:"questionID"`
	UserID       uint   `json:"userID"`
	BookmarkType string `json:"bookmarkType"`
}
