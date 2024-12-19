package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type CourseUsecase interface {
	CreateCourse(data presenter.CourseCreateUpdateRequest) map[string]string
	ListCourse(request presenter.CourseListRequest) ([]presenter.CourseResponse, int, error)
	GetCourseDetails(userID, courseID uint) (*presenter.CourseDetailsPresenter, error)
	UpdateCourse(data presenter.CourseCreateUpdateRequest) map[string]string
	DeleteCourse(id uint) error
	AssignSubjectsToCourse(uint, []uint) map[string]string

	EnrollInCourse(userID, courseID uint) map[string]string
	// SubscribeCourse(request presenter.SubscribeCourseRequest) map[string]string
	ListEnrolledCourse(userID uint, page int, search string, pageSize int) ([]presenter.CourseResponse, int, error)

	UpdateAvailability(id uint) map[string]string
}

type CourseRepository interface {
	GetCourseByCourseID(courseID string) (models.Course, error)
	GetCourseByID(id uint) (models.Course, error)

	CreateCourse(data presenter.CourseCreateUpdateRequest) error
	ListCourse(request presenter.CourseListRequest) ([]models.Course, float64, error)
	GetCourseDetails(courseID uint) (*models.Course, error)
	GetSubjectHeirarchy(id uint) ([]presenter.SubjectHeirarchyDetails, error)
	UpdateCourse(data presenter.CourseCreateUpdateRequest) error
	DeleteCourse(id uint) error

	EnrollInCourse(userID, courseID uint) error
	// SubscribeCourse(request presenter.SubscribeCourseRequest) error
	ListEnrolledCourse(userID uint, page int, search string, pageSize int) ([]models.Course, float64, error)
	GetCourseGroupByCourseID(id uint) ([]models.CourseGroup, error)
	CheckStudentCourse(userID, courseID uint) (models.StudentCourse, error)
	AssignSubjectsToCourse(courseIds []uint, subjectIDs []uint) error
	UpdateAvailability(course models.Course) error
}
