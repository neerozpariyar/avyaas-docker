package presenter

type ReferralListRequest struct {
	Page     int
	PageSize int
	CourseID uint
	Search   string
}

type ReferralListResponse struct {
	Success     bool               `json:"success"`
	CurrentPage int32              `json:"currentPage"`
	TotalPage   int32              `json:"totalPage"`
	Data        []ReferralResponse `json:"data"`
}

type CreateUpdateReferralRequest struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title" validate:"required"`
	Type         string  `json:"type" validate:"required"`
	CourseID     uint    `json:"courseID"`
	UserID       uint    `json:"userID"`
	Code         string  `json:"code" validate:"required"`
	ExpiryDate   string  `json:"expiryDate" validate:"required"`
	DiscountType string  `json:"discountType" validate:"required"`
	Discount     float32 `json:"discount" validate:"required"`
	HasLimit     *bool   `json:"hasLimit" gorm:"default:false"`
	HasUsed      *bool   `json:"hasUsed" gorm:"default:false"`
	Limit        uint    `json:"limit" gorm:"default:0"`
}

type ReferralResponse struct {
	ID           uint        `json:"id"`
	Title        string      `json:"title"`
	Course       interface{} `json:"course"`
	User         interface{} `json:"user"`
	Code         string      `json:"code"`
	Type         string      `json:"type"`
	DiscountType string      `json:"discountType"`
	Discount     float32     `json:"discount"`
	HasLimit     *bool       `json:"hasLimit" gorm:"default:false"`
	HasUsed      *bool       `json:"hasUsed" gorm:"default:false"`
	Limit        uint        `json:"limit"`
	ExpiryDate   string      `json:"expiryDate"`
}

type ApplyReferralRequest struct {
	UserID    uint   `json:"userID"`
	Code      string `json:"code" validate:"required"`
	PackageID uint   `json:"packageID" validate:"required"`
}

type ApplyReferralResponse struct {
	Price           float32 `json:"price"`
	DiscountedPrice float32 `json:"discountedPrice"`
}
