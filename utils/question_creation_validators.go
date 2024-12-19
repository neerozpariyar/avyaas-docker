package utils

import (
	"avyaas/internal/domain/presenter"
	"mime/multipart"
)

func IsValidImageFile(file *multipart.FileHeader) bool {
	fileType := GetFileType(file.Filename)
	return fileType == "png" || fileType == "jpg" || fileType == "jpeg"
}

func IsValidAudioFile(file *multipart.FileHeader) bool {
	fileType := GetFileType(file.Filename)
	return fileType == "mpeg" || fileType == "mp3"
}

// Helper function to count correct options
func CountCorrectOptions(options []presenter.OptionCreate) int {
	count := 0
	for _, option := range options {
		if option.IsCorrect {
			count++
		}
	}
	return count
}

// Helper function to check if options contain media (used in FillInTheBlanks validation)
func HasMediaOptions(options []presenter.OptionCreate) bool {
	for _, option := range options {
		if option.Image != nil || option.Audio != nil {
			return true
		}
	}
	return false
}
