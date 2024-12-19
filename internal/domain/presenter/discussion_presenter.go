package presenter

type DiscussionCreateUpdateRequest struct {
	ID        uint   `json:"id"`
	Title     string `json:"title" validate:"required"`
	Query     string `json:"query" validate:"required"`
	UserID    uint   `json:"userID"`
	SubjectID uint   `json:"subjectID" validate:"required"`
	CourseID  uint   `json:"courseID" validate:"required"`
	CreatedBy uint   `json:"createdBy"`
}
type DiscussionListRequest struct {
	Page         int
	PageSize     int
	SubjectID    uint
	UserID       uint
	DiscussionID uint
	Search       string
}

type Discussion struct {
	ID         uint        `json:"id"`
	Title      string      `json:"title"`
	Query      string      `json:"query"`
	VoteCount  uint        `json:"voteCount"`
	ReplyCount uint        `json:"replyCount"`
	Views      uint        `json:"views"`
	Course     interface{} `json:"course"`
	Subject    interface{} `json:"subject"`
	CreatedBy  interface{} `json:"createdBy"`
	HasLiked   bool        `json:"hasLiked"`
}
