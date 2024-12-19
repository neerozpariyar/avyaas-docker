package usecase

import (
	"avyaas/internal/domain/presenter"
)

func (uCase *usecase) ListStudent(request *presenter.StudentListRequest) ([]presenter.UserResponse, int, error) {
	// Delegate the retrieval of students
	students, totalPage, err := uCase.repo.ListStudent(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	return students, int(totalPage), nil
}
