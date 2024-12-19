package presenter

import (
	"mime/multipart"
)

type CreateUpdateQuestionRequest struct {
	ID                 uint                          `json:"id"`
	Title              string                        `json:"title" validate:"required"`
	Image              *multipart.FileHeader         `json:"image"`
	Audio              *multipart.FileHeader         `json:"audio"`
	Options            []OptionCreate                `json:"options" validate:"required"`
	Position           uint                          `json:"position"`
	ForTest            *bool                         `json:"forTest"`
	SubjectID          uint                          `json:"subjectID" validate:"required"`
	QuestionSetID      uint                          `json:"questionSetID"`
	NegativeMark       float64                       `json:"negativeMark"`
	Description        *string                       `json:"description"`
	Type               string                        `json:"type"`
	CaseQuestionID     *uint                         `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
	Questions          []CreateUpdateQuestionRequest `json:"questions"`      // For CaseBased questions, this will contain the nested questions
	IsTrue             *bool                         `json:"isTrue"`
	NestedQuestionType string                        `json:"nestedQuestionType"`
}

//
//	type OptionCreateUpdateRequest struct {
//		ID         uint                  `json:"id"`
//		Title      string                `json:"title"`
//		QuestionID uint                  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"questionID" validate:"required"`
//		Options    []string              `json:"option"`
//		File       *multipart.FileHeader `json:"file"`
//	}
//----------------------------------------------------------------------------/////////
// type CreateQuestionRequest struct {
// 	Text      string                  `json:"text" form:"text" validate:"required"`
// 	Options   []string                `json:"options" form:"options"`
// 	OptionUrl []*multipart.FileHeader `json:"url" form:"url"`
// }
//---------------------------------------------------------------------------///////////////
// type CreateOptionRequest struct {
// 	Text string                `json:"text" form:"text" validate:"required"`
// 	Url  *multipart.FileHeader `json:"url" form:"url"`
// }

// type CreateQuestionRequest struct {
// 	ID         uint                  `json:"id"`
// 	Text       string                `json:"text" form:"text"`
// 	OptionText []string              `json:"optionText" form:"optionText"`
// 	OptionUrl  *multipart.FileHeader `json:"url" form:"url"`
// }

// type CreateOptionRequest struct {
// 	Text string                `json:"text" form:"text" validate:"required"`
// 	Url  *multipart.FileHeader `json:"url" form:"url"`
// }

type ListQuestionRequest struct {
	PageSize      int
	Page          int
	CourseID      uint
	SubjectID     uint
	QuestionSetID uint
	RequesterID   uint
	Search        string
}

type QuestionListResponse struct {
	ID                 uint                   `json:"id"`
	Title              string                 `json:"title"`
	Description        string                 `json:"description"`
	Image              string                 `json:"image"`
	ForTest            *bool                  `json:"forTest"`
	Courses            []CourseDataForPoll    `json:"courses"`
	Subject            interface{}            `json:"subject"`
	Type               string                 `json:"type"`           // CaseBased, MultiAnswer, FillInTheBlanks, MCQ, TrueOrFalse
	CaseQuestionID     uint                   `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
	NegativeMark       float64                `json:"negativeMark"`   // If this question has negative marks
	IsBookmarked       bool                   `json:"isBookmarked"`
	NestedQuestionType string                 `json:"nestedQuestionType"`
	Options            []OptionListPresenter  `json:"options"`
	QuestionSetID      uint                   `json:"questionSetID"`
	BookmarkID         uint                   `json:"bookmarkID"`
	IsTrue             *bool                  `json:"isTrue"`
	Questions          []QuestionListResponse `json:"questions"` // For CaseBased questions, this will contain the nested questions
	Audio              string                 `json:"audio"`
}

type OptionListPresenter struct {
	ID         uint    `json:"id"`
	QuestionID uint    `json:"questionID"`
	Image      *string `json:"image"`
	Audio      *string `json:"audio"`
	Text       string  `json:"text"`
	IsCorrect  *bool   `json:"isCorrect"`
}
type QuestionPresenter struct {
	ID           uint        `json:"id"`
	Title        string      `json:"title"`
	Image        string      `json:"image"`
	Audio        string      `json:"audio"`
	ForTest      *bool       `json:"forTest"`
	Course       interface{} `json:"course"`
	Subject      interface{} `json:"subject"`
	Options      []Option    `json:"options"`
	Answer       string      `json:"answer"`
	IsBookmarked bool        `json:"isBookmarked"`
	BookmarkID   uint        `json:"bookmarkID"`
}

type OptionCreate struct {
	// Title string                `json:"title" validate:"min=4"`
	// Image *multipart.FileHeader `json:"image"`
	// Audio *multipart.FileHeader `json:"audio"`
	// // File  *multipart.FileHeader `json:"file" validate:"min=4"`
	ID         uint                  `json:"id"`
	QuestionID uint                  `json:"questionID"`
	Image      *multipart.FileHeader `json:"image"`
	Audio      *multipart.FileHeader `json:"audio"`
	Text       string                `json:"text"`
	IsCorrect  bool                  `json:"isCorrect"` //
}

type Option struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Audio string `json:"audio"`
}
