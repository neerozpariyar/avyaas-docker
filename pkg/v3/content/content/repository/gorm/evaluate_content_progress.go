package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
	"math"
)

func (repo *Repository) EvaluateContentProgress(data presenter.ContentProgressPresenter) error {
	userID := data.UserID
	contentID := data.ID
	elapsedDuration := data.ElapsedDuration

	var content models.Content
	if err := repo.db.Model(&models.Content{}).Where("id = ?", contentID).First(&content).Error; err != nil {
		return err
	}

	var cProgress *models.StudentContent

	if content.ContentType == "VIDEO" {
		totalDuration := content.Length

		if elapsedDuration > totalDuration {
			return fmt.Errorf("elapsed Duration: %v can't be greater than the totalDuration: %v", elapsedDuration, totalDuration)
		}

		percentage := float64(0)
		if totalDuration > 0 {
			percentageExact := float64(elapsedDuration) / float64(totalDuration) * 100
			percentage = math.Ceil(percentageExact*100) / 100
		} else if totalDuration == 0 {
			percentage = 100
		}

		cProgress = &models.StudentContent{
			UserID:    userID,
			ContentID: contentID,
			Progress:  float64(percentage),
		}
		if percentage == 100 {
			hasCompleted := true
			cProgress.HasCompleted = &hasCompleted
		} else {
			hasCompleted := false
			cProgress.HasCompleted = &hasCompleted
		}
	} else if content.ContentType == "PDF" {
		if elapsedDuration > 100 || elapsedDuration < 100 {
			elapsedDuration = 100 // as percentage can't be greater than 100
		}
		hasCompleted := true

		cProgress = &models.StudentContent{
			UserID:       userID,
			ContentID:    contentID,
			Progress:     float64(elapsedDuration),
			HasCompleted: &hasCompleted,
		}
	} else {
		return fmt.Errorf("invalid content type: %v", content.ContentType)
	}

	// SaveOrUpdateContentProgress
	if err := repo.SaveOrUpdateContentProgress(cProgress); err != nil {
		return err
	}

	// Get the course ID associated with the content ID
	courseID, err := repo.GetCourseIDByContentID(contentID)
	if err != nil {
		return err
	}

	// Update hasCompleted and progress in StudentContent based on percentage
	hasCompleted := cProgress.Progress == 100
	progress := cProgress.Progress

	if err := repo.db.Model(&models.StudentContent{}).
		Where("user_id = ? AND content_id = ?", userID, contentID).
		Updates(map[string]interface{}{
			"has_completed": hasCompleted,
			"progress":      progress,
		}).Error; err != nil {
		return err
	}

	// Get the actual paid and expiryDate from the StudentContent table
	var studentContent models.StudentContent
	if err := repo.db.Model(&models.StudentContent{}).
		Where("user_id = ? AND content_id = ?", userID, contentID).
		First(&studentContent).Error; err != nil {
		return err
	}

	// Populate the paid and expiryDate fields
	cProgress.Paid = studentContent.Paid
	cProgress.ExpiryDate = studentContent.ExpiryDate

	// Get course progress and update StudentContent with courseID and course progress
	err = repo.EvaluateProgress(presenter.ProgressPresenter{
		UserID:   userID,
		CourseID: courseID,
	})
	if err != nil {
		return err
	}

	// Update courseID and course progress in StudentContent
	cProgress.CourseID = courseID
	// cProgress.CourseProgress = courseProgress.Progress
	cProgress.ID = studentContent.ID

	return nil
}
