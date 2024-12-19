package presenter

type BookmarkCreateUpdateRequest struct {
	ID         uint `json:"id"`
	ContentID  uint `json:"contentID"`
	QuestionID uint `json:"questionID"`
	UserID     uint `json:"userID"`
}

type BookmarkListRequest struct {
	Page         int
	PageSize     int
	BookmarkType string `json:"bookmarkType"`
	UserID       uint   `json:"userID"`
	Search       string
}

type BookmarkListResponse struct {
	ID           uint   `json:"id"`
	ContentID    uint   `json:"contentID"`
	QuestionID   uint   `json:"questionID"`
	UserID       uint   `json:"userID"`
	BookmarkType string `json:"bookmarkType"`
	Title        string `json:"title"`
}

type BookmarkDetailResponse struct {
	ID        uint        `json:"id"`
	ContentID uint        `json:"contentID"`
	Content   interface{} `json:"content"`
	// QuestionSetID uint        `json:"questionSetID"`
	QuestionID uint        `json:"questionID"`
	Question   interface{} `json:"question"`
}
