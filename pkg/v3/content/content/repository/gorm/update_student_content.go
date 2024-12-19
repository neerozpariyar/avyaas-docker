package gorm

func (repo *Repository) UpdateStudentContent(contentID uint) error {
	// transaction := repo.db.Begin()

	// Fetch the corresponding course ID based on content ID
	// var err error
	// var chapterContent models.ChapterContent
	// if err = repo.db.
	// 	Where("content_id = ?", contentID).
	// 	First(&chapterContent).Error; err != nil {
	// 	return err
	// }

	// var chapter models.Chapter
	// if err := repo.db.
	// 	Where("id = ?", chapterContent.ChapterID).
	// 	First(&chapter).Error; err != nil {
	// 	return err
	// }

	// var unit models.Unit
	// if err := repo.db.
	// 	Where("id = ?", chapter.UnitID).
	// 	First(&unit).Error; err != nil {
	// 	return err
	// }

	// var subject models.Subject
	// if err := repo.db.
	// 	Where("id = ?", unit.SubjectID).
	// 	First(&subject).Error; err != nil {
	// 	return err
	// }

	// var course models.Course
	// if err := repo.db.
	// 	Where("id = ?", subject.CourseID).
	// 	First(&course).Error; err != nil {
	// 	return err
	// }

	// var uniqueUserIDs []uint
	// if err := repo.db.
	// 	Model(&models.StudentContent{}).
	// 	Where("course_id = ?", course.ID).
	// 	Distinct("user_id").
	// 	Pluck("user_id", &uniqueUserIDs).Error; err != nil {
	// 	return err
	// }

	// if len(uniqueUserIDs) == 0 {
	// 	if err := repo.db.
	// 		Model(&models.StudentContent{}).
	// 		Where("content_id = ?", contentID).
	// 		Distinct("user_id").
	// 		Pluck("user_id", &uniqueUserIDs).Error; err != nil {
	// 		return err
	// 	}
	// }
	// for _, userID := range uniqueUserIDs {
	// 	var existingStudentContent models.StudentContent
	// 	err := repo.db.
	// 		Where("user_id = ? AND content_id = ? AND course_id = ?", userID, contentID, course.ID).
	// 		First(&existingStudentContent).Error

	// 	if err != nil && err != gorm.ErrRecordNotFound {
	// 		return err
	// 	}

	// 	if err == gorm.ErrRecordNotFound {
	// 		hasCompleted := false
	// 		paid := true

	// 		// Fetch the existing ExpiryDate for the same course and user
	// 		var existingExpiryDate *time.Time
	// 		var expiryDate sql.NullTime
	// 		if err := repo.db.
	// 			Model(&models.StudentContent{}).
	// 			Where("user_id = ? AND course_id = ?", userID, course.ID).
	// 			Select("expiry_date").
	// 			Order("expiry_date DESC").
	// 			Limit(1).
	// 			Scan(&expiryDate).
	// 			Error; err != nil && err != gorm.ErrRecordNotFound {
	// 			return err
	// 		}
	// 		if expiryDate.Valid {
	// 			existingExpiryDate = &expiryDate.Time
	// 		}

	// 		newStudentContent := &models.StudentContent{
	// 			UserID:       userID,
	// 			CourseID:     course.ID,
	// 			ContentID:    contentID,
	// 			Paid:         &paid,
	// 			ExpiryDate:   existingExpiryDate,
	// 			Progress:     0,
	// 			HasCompleted: &hasCompleted,
	// 		}

	// 		if err := transaction.Create(newStudentContent).Error; err != nil {
	// 			transaction.Rollback()
	// 			return err
	// 		}
	// 	}
	// }

	// // Commit the transaction
	// transaction.Commit()
	return nil
}
