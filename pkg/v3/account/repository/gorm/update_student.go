package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils/file"

	"avyaas/utils"
	"fmt"
	"math/rand"
	"strings"
)

// update profile by self only

func (repo *Repository) UpdateStudent(data presenter.StudentCreateUpdateRequest) error {
	updatedUser := &models.User{
		FirstName:  data.FirstName,
		MiddleName: data.MiddleName,
		LastName:   data.LastName,
		Gender:     models.Gender(data.Gender),
		Email:      data.Email,
		Phone:      data.Phone,
	}

	transaction := repo.db.Begin()

	if data.Image != nil {
		fileData, err := file.UploadFile("student", data.Image)
		if err != nil {
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
	}

	user, err := repo.GetUserByID(data.ID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	if user.Image != "" {
		var uFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", user.Image).First(&uFile).Error
		if err == nil {
			if err = repo.db.Model(models.File{}).Where("id = ?", uFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	var existingUsername models.User
	splitEmail := strings.Split(data.Email, "@")
	username := (strings.ToLower(splitEmail[0]))

	if err := repo.db.Where("username = ? ", username).First(&existingUsername).Error; err == nil {
		num := rand.Intn(100)

		data.Username = fmt.Sprintf("%s%d", username, num)
	}

	err = transaction.Where("id=?", user.ID).Updates(&updatedUser).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()
	return nil
}
