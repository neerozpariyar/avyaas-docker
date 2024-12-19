package presenter

type PackageTypeCreateUpdateRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	ServiceIDs  []uint `json:"serviceIDs" validate:"required"`
}

type PackageTypeListRequest struct {
	PageSize int
	Page     int
	Search   string
}
