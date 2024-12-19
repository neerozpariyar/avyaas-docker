package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
	"math"
)

func (repo *Repository) EvaluateProgress(data presenter.ProgressPresenter) error {
	userID := data.UserID
	courseID := data.CourseID

	totalContentCount, err := repo.GetTotalContentCount(courseID)
	if err != nil {
		return err
	}

	consumedContentCount, err := repo.GetConsumedContentCount(userID, courseID)
	if err != nil {
		return err
	}

	if consumedContentCount > totalContentCount {
		return fmt.Errorf("consumedContentCount: %v can't be greater than the totalContentCount: %v", consumedContentCount, totalContentCount)
	}

	percentage := float64(0)
	if totalContentCount > 0 {
		percentageExact := float64(consumedContentCount) / float64(totalContentCount) * 100
		percentage = (math.Ceil(percentageExact*100) / 100)
	} else if totalContentCount == 0 {
		percentage = 100
	}

	progress := &models.StudentCourse{
		UserID:   userID,
		CourseID: courseID,
		Progress: float64(percentage),
	}

	// Update hasCompleted in StudentCourse based on percentage
	if percentage == 100 {
		hasCompleted := true

		progress.HasCompleted = &hasCompleted

	} else {
		hasCompleted := false

		progress.HasCompleted = &hasCompleted

	}
	// SaveOrUpdateProgress
	err = repo.SaveOrUpdateProgress(progress)
	if err != nil {
		return err
	}

	return nil
}
