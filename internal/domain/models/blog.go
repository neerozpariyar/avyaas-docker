package models

type Blog struct {
	Timestamp
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Cover       string      `json:"cover"`
	CourseID    uint        `json:"courseID"`
	SubjectID   uint        `json:"subjectID"`
	Tags        string      `json:"tags"`
	Views       uint        `json:"views"`
	CreatedBy   uint        `json:"createdBy"`
	Likes       uint        `json:"likes"`
	Comment     BlogComment `json:"blogComment"`
}
