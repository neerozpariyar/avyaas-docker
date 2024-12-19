package gorm

import (
	"avyaas/internal/domain/models"

	authority "github.com/Ayata-Incorporation/roles_and_permission/cmd/roles_and_permissions"
)

/*
DeleteTeacher is a repository method responsible for performing deletion of a teacher from the database.

Parameters:
  - repo: A pointer to the repository struct, representing the data access layer for user-related
    operations. It provides access to the underlying database for deleting a teacher.
  - id: An unsigned integer representing the ID of the teacher to be deleted.

Returns:
  - error: An error indicating any issues encountered during the hard deletion of the teacher.
    A nil error signifies a successful deletion.
*/
func (repo *Repository) DeleteTeacher(id uint) error {
	transaction := repo.db.Begin()
	// Create a new instance of the authority.Authority struct
	auth := authority.New(authority.Options{DB: transaction})

	user, err := repo.GetUserByID(id)
	if err != nil {
		transaction.Rollback()
		return err
	}
	if user.Image != "" {
		var uiFile models.File

		err = repo.db.Model(&models.File{}).Where("url = ?", user.Image).First(&uiFile).Error
		if err == nil {
			if err = transaction.Model(models.File{}).Where("id = ?", uiFile.ID).Update("is_active", false).Error; err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	// Perform a hard delete of the user instance with the given ID using the GORM Unscoped method
	err = repo.db.Unscoped().Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the teacher with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("id = ?", id).Delete(&models.Teacher{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Perform a hard delete of the teacher with the given ID using the GORM Unscoped method
	err = transaction.Unscoped().Where("user_id = ?", id).Delete(&models.TeacherSubject{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Revoke user role i.e. delete user and role relations from database
	err = auth.RevokeRole(id, 3)
	if err != nil {
		transaction.Rollback()
		return err
	}

	// Retrieve the assigned permissions of the user
	permissions, err := auth.GetUserPermissions(id)
	if err != nil {
		transaction.Rollback()
		return err
	}
	// Revoke/delete all the permission relation of the user if any
	if len(permissions) != 0 {
		for _, permission := range permissions {
			// Revoke the permission from the role
			err := repo.db.Exec("DELETE FROM user_permission_ints WHERE user_id = ? AND permission_id = ?", user.ID, permission.ID).Error
			if err != nil {
				return err
			}
			// if err = auth.RevokePermission(id, permission.Name); err != nil {
			// 	transaction.Rollback()
			// 	fmt.Printf("err: %v\n", err)
			// 	return err
			// }
		}
	}

	// Commit the changes to the database if no errors found.
	transaction.Commit()

	return nil
}
