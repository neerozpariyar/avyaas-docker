package presenter

import (
	"mime/multipart"
)

type CourseCreateUpdateRequest struct {
	ID             uint                  `json:"id"`
	Title          string                `json:"title" validate:"required"`
	CourseID       string                `json:"courseID" validate:"required"`
	Description    string                `json:"description"`
	Available      bool                  `json:"available"`
	CourseGroupIDs []uint                `json:"courseGroupIDs" validate:"required"`
	File           *multipart.FileHeader `json:"file"`
}

type CourseListRequest struct {
	UserID        uint
	CourseGroupID uint
	Search        string
	Page          int
	PageSize      int
}

type CourseResponse struct {
	ID           uint                   `json:"id"`
	CourseID     string                 `json:"courseID"` // slug of course
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Available    *bool                  `json:"available" gorm:"default:false"`
	Thumbnail    string                 `json:"thumbnail"`
	CourseGroups []CourseGroupForCourse `json:"courseGroups"`
	Progress     float64                `json:"progress"`
	ExpiryDate   string                 `json:"expiryDate"`
}

type CourseGroupForCourse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	CourseGroupId string `json:"courseGroupID"`
}

type AssignSubjectsToCourse struct {
	SubjectIDs []uint `json:"subjectIDs"`
	CourseID   uint   `json:"courseID"`
}
