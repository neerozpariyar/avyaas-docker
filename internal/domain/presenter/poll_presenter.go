package presenter

type PollCreateUpdateRequest struct {
	ID        uint     `json:"id"`
	Question  string   `json:"question" validate:"required"`
	Options   []string `json:"options" validate:"required"`
	CourseID  uint     `json:"courseID" validate:"required"`
	SubjectID uint     `json:"subjectID" validate:"required"`
	CreatedBy uint     `json:"createdBy"`
}

type PollListRequest struct {
	PageSize     int    `json:"pageSize"`
	Page         int    `json:"page"`
	UserID       uint   `json:"userID"`
	SubjectID    uint   `json:"subjectID"`
	DiscussionID uint   `json:"discussionID"`
	Search       string `json:"search"`
}

type Poll struct {
	ID          uint                     `json:"id"`
	CreatedAt   string                   `json:"createdAt"`
	Question    string                   `json:"question"`
	Options     []map[string]interface{} `json:"options"`
	Courses     []CourseDataForPoll      `json:"courses"`
	Subject     map[string]interface{}   `json:"subject"`
	VotedOption uint                     `json:"votedOption"`
	CreatedBy   interface{}              `json:"createdBy"`
}

type CourseDataForPoll struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	CourseID string `json:"courseID"`
}
type PollVoteRequest struct {
	UserID   uint `json:"userID"`
	PollID   uint `json:"pollID"`
	OptionID uint `json:"optionID" validate:"required"`
}
