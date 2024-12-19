package presenter

type TestSeriesCreateUpdateRequest struct {
	ID            uint   `json:"id"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	NoOfTests     int    `json:"noOfTests" validate:"required"`
	CourseID      uint   `json:"courseID" validate:"required"`
	StartDate     string `json:"startDate" validate:"required"`
	IsPackage     bool   `json:"isPackage"`
	PackageTypeID uint   `json:"packageTypeID"`
	Price         int    `json:"price"`
	Period        int    `json:"period" ` // in days
}

type ListTestSeriesRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	CourseID uint   `json:"courseID"`
	Search   string `json:"search"`
}

type TestSeriesListResponse struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	NoOfTests   int         `json:"noOfTests"`
	Course      interface{} `json:"course"`
	StartDate   string      `json:"startDate"`
	IsPackage   bool        `json:"isPackage"`
	PackageType interface{} `json:"packageType"`
	Price       int         `json:"price"`
	Period      int         `json:"period" ` // in days
}
