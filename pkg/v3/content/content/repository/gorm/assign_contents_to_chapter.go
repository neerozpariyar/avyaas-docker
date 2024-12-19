package gorm

import (
	"avyaas/internal/domain/models"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (repo *Repository) AssignContentsToChapter(data models.ChapterContent) error {
	transaction := repo.db.Begin()

	var existingContent []models.ChapterContent

	err := repo.db.Debug().Where("content_id = ? AND chapter_id = ?", data.ContentID, data.ChapterID).Find(&existingContent).Error
	if err == nil && len(existingContent) != 0 {
		return fmt.Errorf("content for the given chapter already exists")
	}

	var maxPosition sql.NullInt64
	if err := repo.db.Debug().Model(&models.ChapterContent{}).
		Where("chapter_id = ?", data.ChapterID).
		Select("MAX(position)").
		Row().Scan(&maxPosition); err != nil && err != sql.ErrNoRows {
		return err
	}

	var currentPosition uint
	if maxPosition.Valid {
		currentPosition = uint(maxPosition.Int64)
	} else {
		currentPosition = 0 // Set to 0 if NULL
	}

	// Increment the position of each existing content by 1.
	if currentPosition > 0 {
		if err := transaction.Debug().Model(&models.ChapterContent{}).
			Where("chapter_id = ?", data.ChapterID).
			Update("position", gorm.Expr("position + 1")).Error; err != nil {
			transaction.Rollback()
			return err
		}
	}

	// Check if the content exists
	var content models.Content
	if err := repo.db.First(&content, data.ContentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("content with ID %d does not exist", data.ContentID)
		}
		return err
	}

	// Create the new ChapterContent
	newChapterContent := &models.ChapterContent{
		ChapterID: data.ChapterID,
		ContentID: data.ContentID,
		Position:  1, // Start the position from 1 for LIFO order
	}

	if err := transaction.Debug().Create(&newChapterContent).Error; err != nil {
		transaction.Rollback()
		return err
	}
	// Commit the transaction
	transaction.Commit()

	//Update StudentContent based on ChapterContent
	if err := repo.UpdateStudentContent(data.ContentID); err != nil {
		return err
	}

	return nil
}

//FFFFFFFFFFFIIIIIIIIIIIIIFFFFFFFFFFFFFFOOOOOOOOOOOoo
// func (repo *repository) AssignContentsToChapter(chapterID uint, contentID uint) error {
// 	// Start a transaction
// 	transaction := repo.db.Begin()

// 	// Step 1: Find the maximum position of the existing content associated with the chapter.
// 	var maxPosition sql.NullInt64
// 	if err := transaction.Model(&models.ChapterContent{}).
// 		Where("chapter_id = ?", chapterID).
// 		Select("MAX(position)").
// 		Row().Scan(&maxPosition); err != nil && err != sql.ErrNoRows {
// 		transaction.Rollback()
// 		return err
// 	}

// 	var newPosition uint
// 	if maxPosition.Valid {
// 		newPosition = uint(maxPosition.Int64) + 1 // Increment the max position by 1
// 	} else {
// 		newPosition = 1 // If no content exists, start from position 1
// 	}

// 	// Step 2: Create a new models.ChapterContent record for the content assigned to the chapter with the incremented position.
// 	newChapterContent := &models.ChapterContent{
// 		ChapterID: chapterID,
// 		ContentID: contentID,
// 		Position:  newPosition,
// 	}
// 	if err := transaction.Create(newChapterContent).Error; err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	// Commit the transaction
// 	return transaction.Commit().Error
// }
