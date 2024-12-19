package presenter

type ListLiveGroupRequest struct {
	CourseID uint   `json:"courseID"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Search   string `json:"search"`
}

type LiveGroupListResponse struct {
	ID          uint        `json:"id"`
	Course      interface{} `json:"course"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	// IsPremium   *bool       `json:"isPremium"`
	// Amount      uint        `json:"amount"`
	// ParticipantLimit uint        `json:"participantLimit"`
}
type LiveGroupCreateUpdatePresenter struct {
	CourseID      uint   `json:"courseID" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"` //shown before
	StartDate     string `json:"startDate" validate:"required"`
	IsPackage     bool   `json:"isPackage"`
	PackageTypeID uint   `json:"packageTypeID"`
	Price         int    `json:"price"`
	Period        int    `json:"period" ` // in days

}
