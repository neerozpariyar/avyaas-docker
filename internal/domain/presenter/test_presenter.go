package presenter

/*
CreateUpdateTestRequest is a data structure representing the request format for creating or
updating a test.
*/
type CreateUpdateTestRequest struct {
	ID            uint   `json:"id"`
	Title         string `json:"title" validate:"required"`
	StartTime     string `json:"startTime" validate:"required"`
	EndTime       string `json:"endTime" validate:"required"`
	Duration      int    `json:"duration" validate:"required"`
	ExtraTime     int    `json:"extraTime"`
	TestTypeID    int    `json:"testTypeID" validate:"required"`
	TestSeriesID  uint   `json:"testSeriesID"`
	IsPublic      bool   `json:"isPublic"`
	IsPremium     bool   `json:"isPremium"`
	IsFree        bool   `json:"isFree"`
	IsMock        bool   `json:"isMock" validate:"required"`
	CreatedBy     uint   `json:"createdBy"`
	CourseID      uint   `json:"courseID" validate:"required"`
	QuestionSetID uint   `json:"questionSetID"`
	IsPackage     *bool  `json:"isPackage"`
	Price         int    `json:"price"`
	Period        int    `json:"period"` // in days
}

type AssignQuestionSetToTestRequest struct {
	TestID        uint `json:"testID" validate:"required"`
	QuestionSetID uint `json:"questionSetID" validate:"required"`
}

type ListTestRequest struct {
	UserID     uint   `json:"userID"`
	CourseID   uint   `json:"courseID"`
	TestTypeID uint   `json:"testTypeID"`
	Status     string `json:"status"`
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
}

/*
TestResponse is a data structure representing the response format for a single test data.
*/
type TestResponse struct {
	ID             uint        `json:"id"`
	Title          string      `json:"title"`
	StartTime      string      `json:"startTime"`
	EndTime        string      `json:"endTime"`
	Duration       int         `json:"duration"`
	ExtraTime      int         `json:"extraTime"`
	TotalQuestions int         `json:"totalQuestions"`
	Marks          int         `json:"marks"`
	Price          int         `json:"price"`
	TestType       interface{} `json:"testType"` //to show type directly
	IsPublic       bool        `json:"isPublic"`
	IsPremium      bool        `json:"isPremium"`
	IsFree         bool        `json:"isFree"`
	IsPackage      bool        `json:"isPackage"`
	CreatedBy      uint        `json:"createdBy"`
	Course         interface{} `json:"course"`
	QuestionSet    interface{} `json:"questionSet"`
	TestSeries     interface{} `json:"testSeries"`
}

type LeaderboardResponse struct {
	Test   interface{}               `json:"test"`
	Course interface{}               `json:"course"`
	Users  []LeaderboardUserResponse `json:"users"`
}

type LeaderboardUserResponse struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	RegistrationNumber string  `json:"registrationNumber"`
	Email              string  `json:"email"`
	Phone              string  `json:"phone"`
	Percentage         float64 `json:"percentage"`
	Score              float64 `json:"score"`
	TotalAttempted     int     `json:"totalAttempted"`
	TotalUnattempted   int     `json:"totalUnattempted"`
	TotalCorrect       int     `json:"totalCorrect"`
	TotalWrong         int     `json:"totalWrong"`
}

type TestResultResponse struct {
	ID               uint                          `json:"id"`
	Title            string                        `json:"title"`
	TotalQuestions   int                           `json:"totalQuestions"`
	Marks            int                           `json:"marks"`
	Type             string                        `json:"type"`
	Course           interface{}                   `json:"course"`
	Score            int                           `json:"score"`
	TotalAttempted   int                           `json:"totalAttempted"`
	TotalUnattempted int                           `json:"totalUnattempted"`
	TotalCorrect     int                           `json:"totalCorrect"`
	TotalWrong       int                           `json:"totalWrong"`
	Questions        []TestResultQuestionPresenter `json:"questions"`
}

type TestHistoryRequest struct {
	UserID   uint   `json:"userID"`
	CourseID uint   `json:"courseID"`
	Type     string `json:"type"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type TestHistoryResponse struct {
	ID               uint        `json:"id"`
	Title            string      `json:"title"`
	TotalQuestions   int         `json:"totalQuestions"`
	Marks            int         `json:"marks"`
	Course           interface{} `json:"course"`
	Score            int         `json:"score"`
	TotalAttempted   int         `json:"totalAttempted"`
	TotalUnattempted int         `json:"totalUnattempted"`
	TotalCorrect     int         `json:"totalCorrect"`
	TotalWrong       int         `json:"totalWrong"`
}

type TestResultQuestionPresenter struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Image       string              `json:"image"`
	Options     OptionListPresenter `json:"options"`

	// Options        OptionDetailsPresenter `json:"options"`
	SelectedOptionID uint `json:"selectedOptionID"`
}
