package models

type TypeQuestion struct {
	Timestamp
	Title          string       `json:"title"`
	Image          string       `json:"image"`
	Audio          string       `json:"audio"`
	Description    *string      `json:"description"`
	Type           string       `json:"type"`           // CaseBased, MultiAnswer, FillInBlanks, MCQ, TrueOrFalse
	CaseQuestionID *uint        `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
	ForTest        *bool        `json:"forTest"`
	IsTrue         *bool        `json:"isTrue" `
	SubjectID      uint         `json:"subjectID"`
	Subject        Subject      `json:"-"`
	NegativeMark   *float64     `json:"negativeMark"` // If this question has negative marks
	Options        []TypeOption `gorm:"foreignKey:QuestionID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"options"`
	// QuestionSetID  *uint        `json:"question_set_id"` // If this question is part of a question set
}

type TypeOption struct {
	Timestamp
	QuestionID uint    `gorm:"foreignKey:ID" json:"questionID"`
	Image      *string `json:"image"`
	Audio      *string `json:"audio"`
	Text       *string `json:"text"`
	IsCorrect  bool    `json:"isCorrect"` // For MultiAnswer questions, multiple options can be marked as correct
}

type QuestionSetTypeQuestion struct {
	QuestionSetID  uint `json:"questionSetID"`
	TypeQuestionID uint `json:"typeQuestionID"`
	Position       uint `json:"position"`
}
