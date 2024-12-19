package presenter

import "time"

// type PackageCreateUpdateRequest struct {
// 	ID              uint   `json:"id"`
// 	Title           string `json:"title" validate:"required"`
// 	Description     string `json:"description"`
// 	CourseID        uint   `json:"courseID" validate:"required"`
// 	Price           int    `json:"price" validate:"required"`
// 	Period          int    `json:"period" validate:"required"`
// 	Discount        int    `json:"discount" gorm:"default:0"`
// 	DiscountedPrice int    `json:"discountedPrice"`
// 	ServiceIDs      []uint `json:"serviceIDs"`
// }

type PackageCreateUpdateRequest struct {
	ID            uint   `json:"id"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	PackageTypeID uint   `json:"packageTypeID" validate:"required"`
	CourseID      uint   `json:"courseID" validate:"required"`
	TestSeriesID  uint   `json:"testSeriesID"`
	TestID        uint   `json:"testID"`
	LiveGroupID   uint   `json:"liveGroupID"`
	LiveID        uint   `json:"liveID"`
	Price         int    `json:"price" validate:"required"`
	Period        int    `json:"period" validate:"required"`
}

type SubscribePackageRequest struct {
	UserID          uint       `json:"userID"`
	PackageID       uint       `json:"packageID" validate:"required"`
	PaymentID       uint       `json:"paymentID"` // not needed in payload, will be handled by backend
	PaymentMethod   string     `json:"paymentMethod" validate:"required"`
	Amount          int        `json:"amount" validate:"required"`
	ExpiryDate      *time.Time `json:"expiryDate"`                          // not needed in payload, will be handled by backend
	TransactionID   string     `json:"transactionID" validate:"required"`   // refId for esewa and transaction_id for khalti returned after payment verification
	TransactionUUID string     `json:"transactionUuid" validate:"required"` // pid for esewa and pidx for khalti returned after payment verification
}

type PackageListRequest struct {
	PageSize int
	Page     int
	CourseID uint
	Search   string
}

type PackageListResponse struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	PackageType interface{} `json:"packageType"`
	Course      interface{} `json:"course"`
	TestSeries  interface{} `json:"testSeries"`
	Test        interface{} `json:"testID"`
	LiveGroup   interface{} `json:"liveGroup"`
	Live        interface{} `json:"live"`
	Price       int         `json:"price" validate:"required"`
	Period      int         `json:"period" validate:"required"` // in days
}
