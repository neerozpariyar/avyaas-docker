package presenter

// import "mime/multipart"

// type TypeQuestionPresenter struct {
// 	ID                 uint                    `json:"id"`
// 	Title              string                  `json:"title"`
// 	Image              *multipart.FileHeader   `json:"image"`
// 	Audio              *multipart.FileHeader   `json:"audio"`
// 	Description        *string                 `json:"description"`
// 	Type               string                  `json:"type"`           // CaseBased, MultiAnswer, FillInBlanks, MCQ, TrueOrFalse
// 	CaseQuestionID     *uint                   `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
// 	ForTest            *bool                   `json:"forTest"`
// 	SubjectID          uint                    `json:"subjectID"`
// 	NegativeMark       *float64                `json:"negativeMark"` // If this question has negative marks
// 	Options            []TypeOptionPresenter   `json:"options"`
// 	QuestionSetID      uint                    `json:"questionSetID"`
// 	Questions          []TypeQuestionPresenter `json:"questions"` // For CaseBased questions, this will contain the nested questions
// 	IsTrue             *bool                   `json:"isTrue"`
// 	NestedQuestionType string                  `json:"nestedQuestionType"`
// }
// type TypeOptionPresenter struct {
// 	ID         uint                  `json:"id"`
// 	QuestionID uint                  `json:"questionID"`
// 	Image      *multipart.FileHeader `json:"image"`
// 	Audio      *multipart.FileHeader `json:"audio"`
// 	Text       string                `json:"text"`
// 	IsCorrect  bool                  `json:"isCorrect"` // For MultiAnswer questions, multiple options can be marked as correct
// }

// // type CaseBasedQuestion struct {
// // 	CaseDescription string
// // 	Questions       []TypeQuestionPresenter
// // }

// type TypeQuestionListPresenter struct {
// 	ID                 uint                        `json:"id"`
// 	Title              string                      `json:"title"`
// 	Image              *string                     `json:"image"`
// 	Audio              *string                     `json:"audio"`
// 	Description        *string                     `json:"description"`
// 	Type               string                      `json:"type"`           // CaseBased, MultiAnswer, FillInTheBlanks, MCQ, TrueOrFalse
// 	CaseQuestionID     *uint                       `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
// 	ForTest            *bool                       `json:"forTest"`
// 	SubjectID          uint                        `json:"subjectID"`
// 	NegativeMark       *float64                    `json:"negativeMark"` // If this question has negative marks
// 	Options            []TypeOptionListPresenter   `json:"options"`
// 	QuestionSetID      uint                        `json:"questionSetID"`
// 	Questions          []TypeQuestionListPresenter `json:"questions"` // For CaseBased questions, this will contain the nested questions
// 	IsTrue             *bool                       `json:"isTrue"`
// 	NestedQuestionType string                      `json:"nestedQuestionType"`
// 	IsBookmarked       bool                        `json:"isBookmarked"`
// 	BookmarkID         uint                        `json:"bookmarkID"`
// 	Course             interface{}                 `json:"course"`
// 	Subject            interface{}                 `json:"subject"`
// }
// type TypeOptionListPresenter struct {
// 	ID         uint    `json:"id"`
// 	QuestionID uint    `json:"questionID"`
// 	Image      *string `json:"image"`
// 	Audio      *string `json:"audio"`
// 	Text       string  `json:"text"`
// 	IsCorrect  *bool   `json:"isCorrect"`
// }

// type TypeQuestionListDetailPresenter struct {
// 	ID                 uint                        `json:"id"`
// 	Title              string                      `json:"title"`
// 	Image              *string                     `json:"image"`
// 	Audio              *string                     `json:"audio"`
// 	Description        *string                     `json:"description"`
// 	Type               string                      `json:"type"`           // CaseBased, MultiAnswer, FillInTheBlanks, MCQ, TrueOrFalse
// 	CaseQuestionID     *uint                       `json:"caseQuestionID"` // For CaseBased questions, this will be the same for all questions in a case
// 	ForTest            *bool                       `json:"forTest"`
// 	SubjectID          uint                        `json:"subjectID"`
// 	NegativeMark       *float64                    `json:"negativeMark"` // If this question has negative marks
// 	Options            []TypeOptionListPresenter   `json:"options"`
// 	QuestionSetID      uint                        `json:"questionSetID"`
// 	Questions          []TypeQuestionListPresenter `json:"questions"` // For CaseBased questions, this will contain the nested questions
// 	IsTrue             *bool                       `json:"isTrue"`
// 	NestedQuestionType string                      `json:"nestedQuestionType"`
// 	IsBookmarked       bool                        `json:"isBookmarked"`
// 	BookmarkID         uint                        `json:"bookmarkID"`
// 	Course             interface{}                 `json:"course"`
// 	Subject            interface{}                 `json:"subject"`
// }
