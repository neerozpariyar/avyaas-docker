package usecase

func (uc *usecase) AssignSubjectsToTeacher(userID uint, subjectIDs []uint) map[string]string {
	var errMap = make(map[string]string)
	for _, subjectID := range subjectIDs {
		_, err := uc.subjectRepo.GetSubjectByID(subjectID)
		if err != nil {
			errMap["error"] = err.Error()
			return errMap
		}
	}
	err := uc.repo.AssignSubjectsToTeacher(userID, subjectIDs)
	if err != nil {
		errMap["error"] = err.Error()
		return errMap
	}
	return errMap
}
