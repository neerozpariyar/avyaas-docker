package presenter

import "mime/multipart"

type StudentListRequest struct {
	Page     int
	PageSize int
	CourseID uint
	Search   string
}

type StudentCreateUpdateRequest struct {
	ID         uint                  `json:"id"`
	FirstName  string                `json:"firstname" validate:"required"`
	MiddleName string                `json:"middleName"`
	LastName   string                `json:"lastName" validate:"required"`
	Username   string                `json:"-"`
	Gender     string                `json:"gender" validate:"required"`
	Email      string                `json:"email" validate:"required"`
	Phone      string                `json:"phone" validate:"required"`
	Image      *multipart.FileHeader `json:"image"`
}
