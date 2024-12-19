package usecase

import "fmt"

func (uCase *usecase) AssignQuestionSetToTest(testID, questionSetID uint) error {
	test, err := uCase.repo.GetTestByID(testID)
	if err != nil {
		return err
	}

	questionSet, err := uCase.questionSetRepo.GetQuestionSetByID(questionSetID)
	if err != nil {
		return err
	}

	if test.CourseID != questionSet.CourseID {
		return fmt.Errorf("cannot assign question set '%s' to the test '%s' of different courses", questionSet.Title, test.Title)
	}

	_, err = uCase.repo.GetTestQuestionSet(testID, questionSetID)
	if err == nil {
		return fmt.Errorf("question set '%s' is already assigned to test '%s'", questionSet.Title, test.Title)
	}

	if err := uCase.repo.AssignQuestionSetToTest(testID, questionSetID); err != nil {
		return err
	}

	return nil
}
