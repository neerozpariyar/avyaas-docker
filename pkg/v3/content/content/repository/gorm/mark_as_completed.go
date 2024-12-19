package gorm

func (repo *Repository) MarkAsCompleted(userID, contentID uint) error {
	transaction := repo.db.Begin()

	var err error

	// Update StudentContent
	err = transaction.Exec(`
		UPDATE student_contents
		SET progress = 100, has_completed = true
		WHERE user_id = ? AND content_id = ?
	`, userID, contentID).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Check if this is the last content of the course
	var isLastContent bool
	err = transaction.Raw(`
		SELECT COUNT(*) = SUM(has_completed) 
		FROM student_contents 
		WHERE user_id = ? AND course_id = (
			SELECT course_id FROM student_contents WHERE user_id = ? AND content_id = ?
		)
	`, userID, userID, contentID).Scan(&isLastContent).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// If it's the last content, update the hasCompleted value to true and progress to 100 of the corresponding course
	if isLastContent {
		err = transaction.Exec(`
			UPDATE student_courses
			SET progress = 100, has_completed = true
			WHERE user_id = ? AND course_id = (
				SELECT course_id FROM student_contents WHERE user_id = ? AND content_id = ?
			)
		`, userID, userID, contentID).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	err = transaction.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
