package presenter

import (
	"mime/multipart"
)

type SubjectCreateUpdateRequest struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title" validate:"required"`
	SubjectID   string                `json:"subjectID" validate:"required"`
	Description string                `json:"description"`
	CourseIDs   []uint                `json:"courseIDs" validate:"required"`
	File        *multipart.FileHeader `json:"file"`
}

type SubjectResponse struct {
	ID          uint               `json:"id"`
	SubjectID   string             `json:"subjectID"` // slug of subject
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Thumbnail   string             `json:"thumbnail"`
	Courses     []CourseForSubject `json:"courses"`
}

type CourseForSubject struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	CourseID string `json:"courseID"`
}
type AssignUnitsToSubject struct {
	SubjectID uint   `json:"subjectID"`
	UnitIDs   []uint `json:"unitIDs"`
}

type SubjectHeirarchyDetails struct {
	ID      uint
	Unit    FilterDetail `json:"unit"`
	Chapter FilterDetail `json:"chapter,omitempty"`
	Content FilterDetail `json:"content,omitempty"`
}

type FilterDetail struct {
	ID               uint   `json:"id,omitempty"`
	Title            string `json:"title,omitempty"`
	Description      string `json:"description,omitempty"`
	Thumbnail        string `json:"thumbnail,omitempty"`
	ContentIsPremium *bool  `json:"contentIsPremium,omitempty"`
	ContentType      string `json:"contentType,omitempty"`
	ContentLength    uint   `json:"contentLength,omitempty"`
	Paid             int    `json:"paid,omitempty"`
}

type SubjectHeirarchy struct {
	UnitChapterContentID uint
	UnitID               uint
	UnitTitle            string
	UnitDescription      string
	UnitThumbnail        string
	ChapterID            uint
	ChapterTitle         string
	ContentID            uint
	ContentTitle         string
	ContentIsPremium     *bool
	ContentType          string
	ContentLength        uint
	IsPaid               int
}

type SubjectDetailResponse struct {
	ID               uint
	SubjectID        string
	Title            string
	Description      string
	Thumbnail        string
	SubjectHeirarchy []SubjectHeirarchyDetails `json:"hierarchy"`
}
