package usecase

func (uCase *usecase) AssignQuestionsToQuestionSet(questionSetID uint, questionIDs []uint) error {
	if _, err := uCase.repo.GetQuestionSetByID(questionSetID); err != nil {
		return err
	}

	for _, qID := range questionIDs {
		// if _, err := uCase.questionRepo.GetQuestionByID(qID); err != nil {
		// 	return err
		// }
		if _, err := uCase.questionRepo.GetQuestionByID(qID); err != nil {
			return err
		}
	}

	if err := uCase.repo.AssignQuestionsToQuestionSet(questionSetID, questionIDs); err != nil {
		return err
	}

	return nil
}
