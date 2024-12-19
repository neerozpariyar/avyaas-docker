package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
)

func (repo *Repository) GetSubjectDetails(id, userID uint) (presenter.SubjectDetailResponse, error) {
	var subject models.Subject

	subjectHierarchy, err := repo.GetSubjectHeirarchy(id, userID)

	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	err = repo.db.Model(&models.Subject{}).Where("id  = ?", id).First(&subject).Error

	if err != nil {
		return presenter.SubjectDetailResponse{}, err
	}

	result := presenter.SubjectDetailResponse{
		ID:               subject.ID,
		SubjectID:        subject.SubjectID,
		Description:      subject.Description,
		Title:            subject.Title,
		Thumbnail:        subject.Thumbnail,
		SubjectHeirarchy: subjectHierarchy,
	}

	return result, nil
}
