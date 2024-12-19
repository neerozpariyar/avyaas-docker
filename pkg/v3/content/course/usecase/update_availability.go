package usecase

import "fmt"

func (uCase *usecase) UpdateAvailability(id uint) map[string]string {
	errMap := make(map[string]string)

	// Check if the course with the given ID exists
	course, err := uCase.repo.GetCourseByID(id)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	fmt.Printf("course: %v\n", course)
	// Call the repository to update the course status
	if err = uCase.repo.UpdateAvailability(course); err != nil {
		return errMap
	}

	return errMap
}
