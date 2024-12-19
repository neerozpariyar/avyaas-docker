package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

type SubjectUsecase interface {
	CreateSubject(data presenter.SubjectCreateUpdateRequest) map[string]string
	GetSubjectDetails(userID, subjectID uint) (presenter.SubjectDetailResponse, error)
	ListSubject(page int, courseID uint, search string, pageSize int) ([]presenter.SubjectResponse, int, error)
	UpdateSubject(data presenter.SubjectCreateUpdateRequest) map[string]string
	DeleteSubject(id uint) error
	AssignUnitsToSubject(uint, []uint) map[string]string
	GetSubjectHeirarchy(uint, uint) ([]presenter.SubjectHeirarchyDetails, error)
}

type SubjectRepository interface {
	GetSubjectBySubjectID(SubjectID string) (models.Subject, error)
	GetSubjectByID(id uint) (models.Subject, error)
	CheckUnitInSubjectHeirarchy(subjectId uint) ([]uint, error)
	AssignUnitsToSubject([]uint, []uint) error
	GetCoursesBySubjectId(id uint) ([]models.Course, error)
	GetCourseIDsBySubjectID(id uint) ([]uint, error)
	CreateSubject(data presenter.SubjectCreateUpdateRequest) error
	GetRelationsBySubjectID(id uint) ([]uint, error)
	GetSubjectDetails(id uint, userID uint) (presenter.SubjectDetailResponse, error)
	ListSubject(page int, courseID uint, search string, pageSize int) ([]models.Subject, float64, error)
	UpdateSubject(data presenter.SubjectCreateUpdateRequest) error
	DeleteSubject(id uint) error
	GetSubjectHeirarchy(uint, uint) ([]presenter.SubjectHeirarchyDetails, error)
}
