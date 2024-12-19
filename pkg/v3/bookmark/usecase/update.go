package usecase

// func (uCase *usecase) UpdateBookmark(data presenter.BookmarkCreateUpdateRequest) map[string]string {
// 	var err error
// 	errMap := make(map[string]string)
// 	// Retrieve the existing bookmark  with the provided bookmark's ID
// 	_, err = uCase.repo.GetBookmarkByID(data.ID)
// 	if err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}
// 	if data.ContentID != 0 {

// 		if _, err := uCase.contentRepo.GetContentByID(data.ContentID); err != nil {
// 			errMap["contentID"] = err.Error()
// 			return errMap
// 		}
// 	}
// 	if data.QuestionID != 0 {
// 		if _, err := uCase.questionRepo.GetQuestionByID(data.QuestionID); err != nil {
// 			errMap["questionID"] = err.Error()
// 			return errMap
// 		}
// 	}

// 	// Delegate the update of bookmark
// 	// if err = uCase.repo.UpdateBookmark(data); err != nil {
// 		errMap["error"] = err.Error()
// 		return errMap
// 	}

// 	return errMap
// }
