package models

// Question represents a database model struct for storing information about a question.
//----------------------------------------------------------//////
// type Question struct {
// 	ID      uint     `gorm:"primaryKey"`
// 	Text    string   `json:"text" validate:"required"`
// 	Options []Option `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"options" form:"options" `
// }

//	type Option struct {
//		ID         uint     `gorm:"primaryKey"`
//		QuestionID uint     `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"questionID" form:"questionID"`
//		Question   Question `json:"-"`
//		Text       string   `json:"text" form:"text" ` //validate:"required"
//		Url        string   `json:"url" form:"url"`    // Assuming you store the file URL
//	}
//
// -------------------------------------------------------------------//////////////////
type Question struct {
	Timestamp

	Title          string   `json:"title"`
	Description    *string  `json:"description"`
	Image          string   `json:"image"`
	Audio          string   `json:"audio"`
	ForTest        *bool    `json:"forTest"`
	SubjectID      uint     `json:"subjectID"`
	Subject        Subject  `json:"-"`
	NegativeMark   float64  `json:"negativeMark"`
	CaseQuestionID *uint    `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
	IsTrue         *bool    `json:"isTrue" `
	Options        []Option `gorm:"foreignKey:QuestionID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"options"`
	Type           string   `json:"type"` // CaseBased, MultiAnswer, FillInBlanks, MCQ, TrueOrFalse
	// OptionID    uint    `json:"option"`
	// Position    uint    `json:"position"`
	// Options     []Option `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"options"`
}

type Option struct {
	Timestamp
	QuestionID uint   `gorm:"foreignKey:ID" json:"questionID"`
	Image      string `json:"image"`
	Audio      string `json:"audio"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"isCorrect"`
}
