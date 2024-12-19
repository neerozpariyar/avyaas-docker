package presenter

import "time"

type ContentDetailsPresenter struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title"`
	IsPremium    *bool   `json:"isPremium"`
	ContentType  string  `json:"contentType"`
	Length       uint    `json:"length"`
	Visibility   *bool   `json:"visibility" gorm:"default:false"`
	Views        int     `json:"views,omitempty" gorm:"default:0"`
	Url          string  `json:"url,omitempty"`
	Paid         *bool   `json:"paid"`
	HasCompleted bool    `json:"hasCompleted"`
	Progress     float64 `json:"progress"`
}

type ChapterDetailsPresenter struct {
	ID      uint                      `json:"id"`
	Title   string                    `json:"title"`
	UnitID  uint                      `json:"-"`
	Content []ContentDetailsPresenter `json:"content"`
}

type UnitDetailsPresenter struct {
	ID          uint                      `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Thumbnail   string                    `json:"thumbnail"`
	SubjectID   uint                      `json:"-"`
	Chapter     []ChapterDetailsPresenter `json:"chapter"`
}

type SubjectDetailsPresenter struct {
	ID          uint                   `json:"id"`
	SubjectID   string                 `json:"subjectID"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Thumbnail   string                 `json:"thumbnail"`
	CourseID    uint                   `json:"-"`
	Unit        []UnitDetailsPresenter `json:"unit"`
}

type CourseDetailsPresenter struct {
	Subject    []SubjectDetailsPresenter `json:"subject"`
	Progress   float64                   `json:"progress"`
	ExpiryDate string                    `json:"expiryDate"`
}

type OptionDetailsPresenter struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	OptionA  string `json:"optionA"`
	OptionB  string `json:"optionB"`
	OptionC  string `json:"optionC"`
	OptionD  string `json:"optionD"`
	ImageA   string `json:"imageA"`
	ImageB   string `json:"imageB"`
	ImageC   string `json:"imageC"`
	ImageD   string `json:"imageD"`
	AudioA   string `json:"audioA"`
	AudioB   string `json:"audioB"`
	AudioC   string `json:"audioC"`
	AudioD   string `json:"audioD"`
	Answer   string `json:"answer"`
	FileType string `json:"fileType"`
}

type QuestionDetailsPresenter struct {
	ID          uint                   `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Position    uint                   `json:"position"`
	ForTest     *bool                  `json:"forTest"`
	Subject     interface{}            `json:"subject"`
	Options     OptionDetailsPresenter `json:"options"`
	// SubjectID   uint                   `json:"subjectID"`
	// Options     []OptionDetailsPresenter `json:"options"`
}

type QuestionSetDetailsPresenter struct {
	ID             uint          `json:"id"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	TotalQuestions int           `json:"totalQuestions"`
	Marks          int           `json:"marks"`
	Course         interface{}   `json:"course"`
	File           string        `json:"file"`
	Questions      []interface{} `json:"questions"`
}

type TestDetailsPresenter struct {
	ID             uint                         `json:"id"`
	Title          string                       `json:"title"`
	StartTime      *time.Time                   `json:"startTime"`
	EndTime        *time.Time                   `json:"endTime"`
	Duration       int                          `json:"duration"`
	ExtraTime      int                          `json:"extraTime"`
	TotalQuestions int                          `json:"totalQuestions"`
	Marks          int                          `json:"marks"`
	Price          int                          `json:"price"`
	TestType       interface{}                  `json:"type"`
	IsPublic       *bool                        `json:"isPublic"`
	IsPremium      *bool                        `json:"isPremium"`
	CreatedBy      uint                         `json:"createdBy"`
	Course         interface{}                  `json:"course"`
	QuestionSet    *QuestionSetDetailsPresenter `json:"questionSet"`

	// QuestionSetID uint                          `json:"questionSetID"`
	// QuestionSets  []QuestionSetDetailsPresenter `json:"questionSets"`
}

type SubmitQuestionRequest struct {
	QuestionID uint `json:"id"`
	AnswerID   uint `json:"answerID"`
}

type SubmitTestRequest struct {
	UserID    uint                    `json:"userID"`
	TestID    uint                    `json:"testID"`
	Questions []SubmitQuestionRequest `json:"questions"`
}
