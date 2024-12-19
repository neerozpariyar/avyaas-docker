package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils"
	"avyaas/utils/file"
	"errors"

	"strings"
)

func (repo *Repository) CreateContent(data presenter.ContentCreateUpdateRequest) error {
	content := &models.Content{
		Title:       data.Title,
		Description: data.Description,
		IsPremium:   data.IsPremium,
		ContentType: strings.ToUpper(data.ContentType),
		Length:      data.Length,
		Level:       data.Level,
		Visibility:  data.Visibility,
		CreatedBy:   data.CreatedBy,
	}

	transaction := repo.db.Begin()

	if data.File != nil {
		fileData, err := file.UploadFile("content", data.File)
		if err != nil {
			transaction.Rollback()
			return err
		}

		isActive := true
		urlObject := utils.GetURLObject(fileData.Url)

		err = transaction.Create(&models.File{
			Title:    fileData.Filename,
			Type:     fileData.FileType,
			Url:      urlObject,
			IsActive: &isActive,
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		// rData, err := data.File.Open()
		// if err != nil {
		// 	return err
		// }

		// if strings.ToUpper(data.ContentType) == "VIDEO" {
		// 	// Note: change the fle length check package later
		// 	pData, err := ffprobe.ProbeReader(context.Background(), rData)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	content.Length = uint(pData.Format.DurationSeconds)
		// }

		content.Url = urlObject
	}

	// var existingContents []models.Content
	// repo.db.Preload("Contents").First(&models.Chapter{})
	// // var chapterContents models.ChapterContent
	// err = repo.db.Model(&models.Content{}).Where("chapter_id = ?", data.ChapterID).Find(&existingContents).Order("position").Error
	// if err != nil {
	// 	transaction.Rollback()
	// 	return err
	// }

	var newContent models.Content
	if data == (presenter.ContentCreateUpdateRequest{}) {
		return errors.New("data is nil")
	}

	err := transaction.Create(&content).Scan(&newContent).Error

	if newContent == (models.Content{}) {
		return errors.New("newContent is nil")
	}

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	if data.HasNote != nil && *data.HasNote {
		if data.Note == (presenter.ContentNoteCreateUpdateRequest{}) {
			return errors.New("data.Note is nil")
		}
		noteData := presenter.NoteCreateUpdateRequest{
			Title:       data.Note.Title,
			Description: data.Note.Description,
			ContentID:   newContent.ID,
			File:        data.Note.File,
		}
		data.Note.ContentID = newContent.ID

		err = repo.noteRepo.CreateNote(noteData)
		if err != nil {
			// transaction.Rollback()
			return err
		}
	}
	// Retrieve the ID of the created content
	// transaction.Commit()

	if data.ChapterID != 0 {
		if err := repo.AssignContentsToChapter(models.ChapterContent{
			ChapterID: data.ChapterID,
			ContentID: newContent.ID,
		}); err != nil {
			// transaction.Rollback()
			return err
		}
	}
	// if err == nil && len(existingContents) != 0 {
	// 	err = transaction.Model(models.Content{}).Where("id = ?", newContent.ID).Update("position", existingContents[len(existingContents)-1].Position+1).Error

	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// } else {
	// 	err = transaction.Model(models.Content{}).Where("id = ?", newContent.ID).Update("position", 1).Error

	// 	if err != nil {
	// 		transaction.Rollback()
	// 		return err
	// 	}
	// }

	// transaction.Commit()
	return err
}
