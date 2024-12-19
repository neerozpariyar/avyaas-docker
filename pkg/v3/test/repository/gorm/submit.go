package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func (repo *Repository) SubmitTest(data presenter.SubmitTestRequest) error {
	transaction := repo.db.Begin()

	test, err := repo.GetTestByID(data.TestID)
	if err != nil {
		return err
	}

	testType, err := repo.GetTestTypeByID(uint(test.TestTypeID))
	if err != nil {
		return err
	}

	var questionSet *models.QuestionSet
	err = repo.db.Where("id = ?", test.QuestionSets[0].ID).First(&questionSet).Error
	if err != nil {
		return err
	}

	eachMark := questionSet.Marks / questionSet.TotalQuestions

	totalCorrect := 0
	totalWrong := 0
	totalMarks := 0.0

	for _, question := range data.Questions {

		qs, err := repo.questionRepo.GetQuestionByID(question.QuestionID)
		if err != nil {
			return err
		}

		correctOptionId, err := repo.questionRepo.GetCorrectAnswersIDForQuestion(question.QuestionID)
		if err != nil {
			return errors.New("no option found for question")
		}

		if correctOptionId == question.AnswerID {
			totalMarks += float64(eachMark)
			totalCorrect++
		} else {
			totalMarks -= float64(qs.NegativeMark) // Subtract negativeMarks for wrong answers
			totalWrong++
		}

		err = transaction.Create(&models.TestResponse{
			UserID:     data.UserID,
			TestID:     data.TestID,
			QuestionID: question.QuestionID,
			AnswerID:   question.AnswerID,
		}).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	totalAttempted := totalCorrect + totalWrong
	percentage := float64(totalCorrect) / float64(questionSet.TotalQuestions) * 100

	err = transaction.Create(&models.TestResult{
		UserID:           data.UserID,
		CourseID:         test.CourseID,
		TestID:           data.TestID,
		Type:             testType.Title,
		StartTime:        test.StartTime,
		EndTime:          test.EndTime,
		Score:            float64(totalMarks),
		Percentage:       percentage,
		TotalAttempted:   totalAttempted,
		TotalUnattempted: questionSet.TotalQuestions - totalAttempted,
		TotalCorrect:     totalCorrect,
		TotalWrong:       totalWrong,
	}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	sTest, err := repo.GetStudentTest(data.UserID, data.TestID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Get the number of attempts allowed for the mock test
	nAttempts := viper.GetUint("test_attempts")

	if sTest != nil {
		if sTest.Attempt == nAttempts-1 {
			err = transaction.Debug().Model(&models.StudentTest{}).Where("user_id = ? AND test_id = ?", sTest.UserID, sTest.TestID).UpdateColumn("has_attended", true).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	// If the test is a free type, it won't have StudentTest instance, hence create it.
	// But the paid and premium type test will have StudentTest instance create during subscription, hence update it.
	if *test.IsFree {

		if sTest == nil {
			sTest = &models.StudentTest{
				UserID:      data.UserID,
				CourseID:    test.CourseID,
				TestID:      data.TestID,
				HasAttended: false,
				Attempt:     1,
			}
			if test.IsMock != nil && !*test.IsMock {
				sTest.HasAttended = true
			}
			err = transaction.Create(sTest).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		if sTest != nil {
			// Check if the test is a mock test and the attempts are less than nAttempts
			if *test.IsMock && sTest.Attempt < nAttempts {
				// Increment the Attempt field
				err = transaction.Model(&models.StudentTest{}).
					Where("user_id = ? AND test_id = ?", sTest.UserID, sTest.TestID).
					Update("attempt", sTest.Attempt+1).Error
				if err != nil {
					transaction.Rollback()
					return err
				}
			}

			if sTest.HasAttended && sTest.Attempt >= nAttempts {

				// If the test is a mock test and the attempts are nAttempts or more, return an error

				return fmt.Errorf("you can't attempt this test more than %v times", nAttempts)
			}
		}
	} else {

		// Check if the test is a mock test and the attempts are less than nAttempts
		if *test.IsMock && sTest.Attempt < nAttempts {
			// Increment the Attempt field
			err = transaction.Model(&models.StudentTest{}).
				Where("user_id = ? AND test_id = ?", sTest.UserID, sTest.TestID).
				Update("Attempt", sTest.Attempt+1).Error
			if err != nil {
				transaction.Rollback()
				return err
			}
		}

		if *test.IsMock && sTest.Attempt >= nAttempts {
			// If the test is a mock test and the attempts are nAttempts or more, return an error
			return fmt.Errorf("you can't attempt this test more than %v times", nAttempts)
		}
	}

	transaction.Commit()
	return nil
}
