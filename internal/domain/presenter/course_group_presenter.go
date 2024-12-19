package presenter

import (
	"mime/multipart"
)

/*
CourseGroupCreateUpdateRequest represents the structure for handling course group creation or update
requests.
*/
type CourseGroupCreateUpdateRequest struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title" validate:"required"`
	GroupID     string                `json:"groupID" validate:"required"`
	Description string                `json:"description"`
	File        *multipart.FileHeader `json:"file"`
}

/*
CourseGroupListResponse represents a response structure for listing course groups. It includes
information about the success status, current page, total page count, and a slice of CourseGroup
model, containing details about individual course groups.
*/
type CourseGroupListResponse struct {
	ID          uint   `json:"id"`
	GroupID     string `json:"groupID"` // slug of course group
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	NoOfCourses int    `json:"noOfCourses"`
}

type AssignCoursesToCourseGroup struct {
	CourseGroupID uint   `json:"courseGroupID"`
	CourseIDs     []uint `json:"courseIDs"`
}
