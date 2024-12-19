package presenter

import (
	"mime/multipart"
)

type UnitCreateUpdateRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	SubjectIDs  []uint `json:"subjectIDs" validate:"required"`
	Position    uint
	File        *multipart.FileHeader `json:"file"`
}

type UpdateUnitPositionRequest struct {
	UnitIDs []uint `json:"unitIDs"`
}

type Unit struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Thumbnail   string             `json:"thumbnail"`
	Course      []CourseForUnit    `json:"courses"`
	Subject     []SubjectForCourse `json:"subjects"`
}

// Title        string `json:"title"`
// Description string `json:"description"`
// Thumbnail   string `json:"thumbnail"`
// UnitID   uint   `json:"subjectID"`
// Subject     Subject
// Chapters    []Chapter `json:"chapters"`
type AssignChaptersToUnit struct {
	ChapterIDs []uint `json:"chapterIDs"`
	RelationID uint   `json:"relationID"`
	SubjectID  uint   `json:"subjectID"`
}

type CourseForUnit struct {
	ID       uint
	Title    string
	CourseID string
}

type SubjectForCourse struct {
	ID        uint
	Title     string
	SubjectID string
}
