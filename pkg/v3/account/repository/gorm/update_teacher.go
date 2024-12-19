package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils/file"

	"avyaas/utils"

	"gorm.io/gorm"
)

/*
UpdateTeacher is a repository method responsible for updating teacher information in the database.

Parameters:
  - repo: A pointer to the repository struct, representing the data access layer for user-related
    operations. It provides access to the underlying database for updating teacher information.
  - data: An instance of the TeacherCreateUpdateRequest struct containing the necessary information
    for updating a teacher.

Returns:
  - error: An error indicating any issues encountered during the update of teacher information.
    A nil error signifies a successful update.
*/
func (repo *Repository) UpdateTeacher(data presenter.TeacherCreateUpdateRequest) error {
	var err error
	var user models.User

	updatedUser := &models.User{
		FirstName:  data.FirstName,
		MiddleName: data.MiddleName,
		LastName:   data.LastName,
		Username:   data.Username,
		Gender:     models.Gender(data.Gender),
		Email:      data.Email,
		Phone:      data.Phone,
	}

	transaction := repo.db.Begin()

	if data.Image != nil {
		fileData, err := file.UploadFile("teacher", data.Image)
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

		updatedUser.Image = urlObject
	}

	// Initiate the user instance update
	if err = transaction.Omit("created_at").Save(&models.User{
		Timestamp: models.Timestamp{
			ID: data.ID,
		},
		FirstName:  updatedUser.FirstName,
		MiddleName: updatedUser.MiddleName,
		LastName:   updatedUser.LastName,
		Username:   updatedUser.Username,
		Email:      updatedUser.Email,
		Gender:     updatedUser.Gender,
		Phone:      updatedUser.Phone,
		Image:      updatedUser.Image,
		RoleID:     3,
	}).Scan(&user).Error; err != nil {
		transaction.Rollback()
		return err
	}

	// Update the teacher instance
	for _, subjectID := range data.SubjectIDs {
		// Initiate the creation of teacher instance in the database
		err = repo.UpdateTeacherData(user, subjectID, transaction)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	// Commit the changes to the database if no errors found.
	transaction.Commit()

	return nil
}

/*
UpdateTeacherData is a repository method responsible for updating additional teacher-related data in the database.
It is called within the UpdateTeacher function's database transaction and updates the CourseID and SubjectID fields
of the associated Teacher record based on the provided user information.

Parameters:
  - user: An instance of the User model representing the updated teacher information.
  - courseID: An unsigned integer representing the ID of the course associated with the teacher.
  - subjectID: An unsigned integer representing the ID of the subject associated with the teacher.
  - transaction: A pointer to the GORM database transaction for atomic operations.

Returns:
  - error: An error indicating any issues encountered during the association of the teacher with a course and subject.
    A nil error signifies a successful association.
*/
func (repo *Repository) UpdateTeacherData(user models.User, subjectID uint, transaction *gorm.DB) error {
	// Initiate the update of teacher instance
	var courseID uint
	err := repo.db.Debug().Model(&models.Subject{}).Select("course_id").Where("id= ?", subjectID).Scan(&courseID).Error
	if err != nil {
		return err
	}
	err = transaction.Where("id=?", user.ID).Updates(&models.Teacher{
		Timestamp: models.Timestamp{
			UpdatedAt: user.UpdatedAt,
		},
	}).Error
	if err != nil {
		return err
	}

	return nil
}
