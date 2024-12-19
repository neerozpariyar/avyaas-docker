package models

// type Package struct {
// 	Timestamp

// 	Title           string    `json:"title" validate:"required"`
// 	Description     string    `json:"description"`
// 	CourseID        uint      `json:"courseID" validate:"required"`
// 	Course          Course    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
// 	Price           int       `json:"price" validate:"required"`
// 	Period          int       `json:"period" validate:"required"` // in days
// 	Discount        int       `json:"discount" gorm:"default:0"`
// 	DiscountedPrice int       `json:"discountedPrice"`
// 	Services        []Service `gorm:"many2many:package_services" json:"services"`
// }

type Package struct {
	Timestamp

	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	PackageTypeID uint   `json:"packageTypeID"`
	CourseID      uint   `json:"courseID" validate:"required"`
	TestSeriesID  uint
	TestID        uint
	LiveGroupID   uint
	LiveID        uint
	Price         int `json:"price" validate:"required"`
	Period        int `json:"period" validate:"required"` // in days
}

type PackageType struct {
	Timestamp

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Services    []Service `gorm:"many2many:package_type_services" json:"services"`
}
