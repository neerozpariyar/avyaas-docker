package models

// QuestionSet represents a database model struct for storing information about a set of questions.
type QuestionSet struct {
	Timestamp

	Title          string     `json:"title"`
	Description    string     `json:"description"`
	TotalQuestions int        `json:"totalQuestions"`
	Marks          int        `json:"marks"`
	CourseID       uint       `json:"courseID"`
	Course         Course     `json:"-"`
	Questions      []Question `gorm:"many2many:question_set_questions;" json:"questions"`
	File           string     `json:"file"`
}

type QuestionSetQuestion struct { // for many2many relation only
	QuestionSetID  uint `json:"questionSetID"`
	QuestionID     uint `json:"questionID"`
	TypeQuestionID uint `json:"typeQuestionID"`
	Position       uint `json:"position"`
}
