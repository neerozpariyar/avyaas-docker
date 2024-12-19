package usecase

import (
	"avyaas/internal/domain/presenter"
	"fmt"
)

func (uCase *usecase) SubmitTest(data presenter.SubmitTestRequest) map[string]string {
	var err error
	errMap := make(map[string]string)

	// Check if the test with the given ID exists
	test, err := uCase.repo.GetTestByID(data.TestID)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	user, err := uCase.accountRepo.GetUserByID(data.UserID)
	if err != nil {
		errMap["error"] = fmt.Sprintf("error getting user by id: %v", err)
		return errMap
	}
	if user.RoleID == 4 {

		studentTest, err := uCase.repo.GetStudentTest(data.UserID, data.TestID)
		if err != nil {
			if err.Error() == fmt.Sprintf("association for student id:'%d' with test id: '%d' not found", data.UserID, data.TestID) {

				if !*test.IsFree {
					errMap["error"] = err.Error()
					return errMap
				}

			} else {
				errMap["error"] = err.Error()
				return errMap
			}

		}

		if test.IsMock != nil && !*test.IsMock {
			if studentTest != nil && studentTest.HasAttended {
				errMap["error"] = "already attempted the test"
			}
		}
		// Call the repository to submit the test
		if err = uCase.repo.SubmitTest(data); err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	}
	return errMap
}
