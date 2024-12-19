package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
Usecase represents the account usecase interface, defining methods for handling various account
related operations.
*/
type AccountUsecase interface {
	ChangePassword(request presenter.ChangePasswordRequest) map[string]string

	CreateTeacher(data presenter.TeacherCreateUpdateRequest) map[string]string
	ListTeacher(request *presenter.TeacherListRequest) ([]presenter.TeacherListResponse, int, error)
	UpdateTeacher(data presenter.TeacherCreateUpdateRequest) map[string]string
	DeleteTeacher(id uint) error
	AssignSubjectsToTeacher(userID uint, subjectID []uint) map[string]string
	ListTeacherReferrals(id uint) ([]presenter.TeacherReferralList, error)

	UpdateStudent(data presenter.StudentCreateUpdateRequest) map[string]string
	ListStudent(request *presenter.StudentListRequest) ([]presenter.UserResponse, int, error)
}

/*
Repository represents the account repository interface, defining methods for handling various account
related operations.
*/
type AccountRepository interface {
	GetUserByID(id uint) (*presenter.UserResponse, error)
	GetUserByEmail(email string) (*presenter.UserResponse, error)
	GetUserByUsername(username string) (*presenter.UserResponse, error)
	GetUserByPhone(phone string) (*presenter.UserResponse, error)
	GetTeacherByID(id uint) (*models.Teacher, error)
	GetTeacherByReferralCode(referral string) (*models.Teacher, error)

	ChangePassword(request presenter.ChangePasswordRequest) error

	CreateTeacher(data presenter.TeacherCreateUpdateRequest) error
	ListTeacher(request *presenter.TeacherListRequest) ([]presenter.UserResponse, float64, error)
	UpdateTeacher(data presenter.TeacherCreateUpdateRequest) error
	DeleteTeacher(id uint) error
	AssignSubjectsToTeacher(userID uint, subjectID []uint) error

	GetStudentByID(id uint) (*models.Student, error)
	UpdateStudent(data presenter.StudentCreateUpdateRequest) error
	ListStudent(request *presenter.StudentListRequest) ([]presenter.UserResponse, float64, error)

	GetStudentsByTeacherReferral(id uint) ([]models.Student, error)
	GetSubscriptionByUserID(userID uint) ([]models.Subscription, error)
}
