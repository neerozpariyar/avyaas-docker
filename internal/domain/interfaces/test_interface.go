package interfaces

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

/*
Usecase represents the test usecase interface, defining methods for handling various test related
operations.
*/
type TestUsecase interface {
	CreateTestType(testType models.TestType) map[string]string
	ListTestType(page int, pageSize int) ([]models.TestType, int, error)
	UpdateTestType(testType models.TestType) map[string]string
	DeleteTestType(id uint) error

	CreateTest(data presenter.CreateUpdateTestRequest) map[string]string
	GetTestDetails(testID, requesterID uint) (*presenter.TestDetailsPresenter, error)
	ListTest(request presenter.ListTestRequest) ([]presenter.TestResponse, int, error)
	UpdateTest(data presenter.CreateUpdateTestRequest) map[string]string
	DeleteTest(id uint) error
	AssignQuestionSetToTest(testID, questionSetID uint) error
	UpdateTestStatus(id uint) map[string]string

	SubmitTest(data presenter.SubmitTestRequest) map[string]string
	GetTestLeaderboard(testID uint) (*presenter.LeaderboardResponse, error)
	GetTestResult(testID, requesterID uint) (*presenter.TestResultResponse, error)
	GetTestHistory(request presenter.TestHistoryRequest) ([]presenter.TestHistoryResponse, float64, error)
}

/*
Repository represents the test repository interface, defining methods for handling various test
related operations.
*/
type TestRepository interface {
	GetTestTypeByID(id uint) (models.TestType, error)
	GetTestTypeByName(title string) (models.TestType, error)
	CreateTestType(testType models.TestType) error
	ListTestType(page int, pageSize int) ([]models.TestType, float64, error)
	UpdateTestType(testType models.TestType) error
	DeleteTestType(id uint) error

	GetTestByID(id uint) (models.Test, error)
	GetTestsByTestSeriesID(testSeriesID uint) ([]models.Test, error)
	// GetTestDetails(testID, requesterID uint) (*presenter.TestDetailsPresenter, error)
	CreateTest(data presenter.CreateUpdateTestRequest) map[string]string
	ListTest(request presenter.ListTestRequest) ([]models.Test, float64, error)
	UpdateTest(data presenter.CreateUpdateTestRequest) map[string]string
	DeleteTest(id uint) error
	AssignQuestionSetToTest(testID, questionSetID uint) error
	GetTestQuestionSet(testID, questionSetID uint) (*models.Test, error)
	UpdateTestStatus(test models.Test) error

	SubmitTest(data presenter.SubmitTestRequest) error
	GetTestLeaderboard(testID uint) ([]models.TestResult, error)
	GetTestResult(testID, requesterID uint) ([]models.TestResponse, error)
	GetStudentTest(userID, testID uint) (*models.StudentTest, error)
	GetStudentTestResult(userID, testID uint) (*models.TestResult, error)
	GetTestHistory(request presenter.TestHistoryRequest) ([]models.TestResult, float64, error)
}
