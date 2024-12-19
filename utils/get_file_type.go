package utils

import (
	"avyaas/internal/domain/models"
	"log"
	"strings"

	"gorm.io/gorm"
)

func GetFileType(fileName string) string {
	splitArray := strings.Split(fileName, ".")

	return splitArray[(len(splitArray) - 1)]
}

func UpdateFileIsActive(url string, transaction *gorm.DB) error {
	var file models.File

	err := transaction.Debug().Model(&models.File{}).Where("url = ?", url).First(&file).Error
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if err == nil {
		if err = transaction.Debug().Model(models.File{}).Where("id = ?", file.ID).Update("is_active", false).Error; err != nil {
			log.Printf("%v", err)
			return err
		}
	}

	return err
}
