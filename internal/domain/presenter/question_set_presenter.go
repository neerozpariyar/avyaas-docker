package presenter

import (
	"mime/multipart"
)

/*
CreateUpdateQuestionSetRequest is a data structure representing the request format for creating or
updating a question set.
*/
type CreateUpdateQuestionSetRequest struct {
	ID             uint                          `json:"id"`
	Title          string                        `json:"title" validate:"required"`
	Description    string                        `json:"description"`
	TotalQuestions int                           `json:"totalQuestions" validate:"required"`
	Marks          int                           `json:"marks" validate:"required"`
	CourseID       uint                          `json:"courseID" validate:"required"`
	File           *multipart.FileHeader         `json:"file"`
	Questions      []CreateUpdateQuestionRequest `json:"questions"`
}

type AssignQuestionsToQuestionSetRequest struct {
	QuestionSetID uint   `json:"questionSetID" validate:"required"`
	QuestionIDs   []uint `json:"questionIDs" validate:"required"`
}
