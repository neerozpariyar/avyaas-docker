package presenter

import (
	"mime/multipart"
)

/*
TeacherCreateUpdateRequest represents the data structure used for creating or updating a teacher
entity.
*/
type TeacherCreateUpdateRequest struct {
	ID         uint                  `json:"id"`
	FirstName  string                `json:"firstname" validate:"required"`
	MiddleName string                `json:"middleName"`
	LastName   string                `json:"lastName" validate:"required"`
	Username   string                `json:"-"`
	Gender     string                `json:"gender" validate:"required"`
	Email      string                `json:"email" validate:"required"`
	Phone      string                `json:"phone" validate:"required"`
	RoleID     int                   `json:"-"`
	SubjectIDs []uint                `json:"subjectIDs" validate:"required"`
	Image      *multipart.FileHeader `json:"image"`
}

type TeacherListRequest struct {
	Page          int
	PageSize      int
	CourseID      uint
	SubjectID     uint
	Search        string
	ReferralCount uint
}

type TeacherListResponse struct {
	ID            uint        `json:"id"`
	FirstName     string      `json:"firstName"`
	MiddleName    string      `json:"middleName"`
	LastName      string      `json:"lastName"`
	Email         string      `json:"email"`
	Phone         string      `json:"phone"`
	Gender        string      `json:"gender"`
	Image         string      `json:"image"`
	Course        interface{} `json:"course"`
	Subject       interface{} `json:"subject"`
	ReferralCode  string      `json:"referralCode"`
	ReferralCount uint        `json:"referralCount"`
}
