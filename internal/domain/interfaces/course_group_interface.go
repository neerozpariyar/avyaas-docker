package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
CourseGroupUsecase represents the course group usecase interface, defining methods for handling
various course group related operations.
*/
type CourseGroupUsecase interface {
	CreateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) map[string]string
	ListCourseGroup(page int, search string, pageSize int) ([]presenter.CourseGroupListResponse, int, error)
	UpdateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) map[string]string
	DeleteCourseGroup(id uint) error
	AssignCoursesToCourseGroup(courseGroupID uint, courseIds []uint) map[string]string
}

/*
CourseGroupRepository represents the course group repository interface, defining methods for handling
various course group related operations.
*/
type CourseGroupRepository interface {
	GetCourseGroupByID(id uint) (models.CourseGroup, error)
	GetCourseGroupByGroupID(groupID string) (models.CourseGroup, error)
	AssignCoursesToCourseGroup(courseGroupIds []uint, courseIds []uint) error

	CreateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) error
	ListCourseGroup(page int, search string, pageSize int) ([]models.CourseGroup, float64, error)
	UpdateCourseGroup(data presenter.CourseGroupCreateUpdateRequest) error
	DeleteCourseGroup(id uint) error
}
