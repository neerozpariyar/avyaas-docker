package presenter

type MarkAsCompleted struct {
	UserID    uint `json:"userID"`
	ContentID uint `json:"contentID"`
	// CourseID     uint  `json:"courseID"`
	HasCompleted *bool `json:"hasCompleted"`
	Progress     uint  `json:"progress"`
}
