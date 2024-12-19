package gorm

import (
	"avyaas/internal/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
GetCourseGroupByID is a repository method responsible for retrieving a course group from the database
based on its unique identifier (ID).

Parameters:
  - id: A uint representing the unique identifier (ID) of the course group to be retrieved.

Returns:
  - courseGroup: A models.CourseGroup instance representing the retrieved course group.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetCourseGroupByID(id uint) (models.CourseGroup, error) {
	var courseGroup models.CourseGroup

	// Retrieve the course group from the database based on given id
	err := repo.db.Where("id = ?", id).First(&courseGroup).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.CourseGroup{}, fmt.Errorf("course group with group id: '%d' not found", id)
		}

		return models.CourseGroup{}, err
	}

	return courseGroup, nil
}

/*
GetCourseGroupByGroupID is a repository method responsible for retrieving a course group from the
database based on its GroupID.

Parameters:
  - groupID: A string representing the unique identifier (GroupID) of the course group to be retrieved.

Returns:
  - courseGroup: A models.CourseGroup instance representing the retrieved course group.
  - error:    An error, if any, encountered during the database retrieval operation.
*/
func (repo *Repository) GetCourseGroupByGroupID(groupID string) (models.CourseGroup, error) {
	var courseGroup models.CourseGroup

	// Retrieve the course group from the database based on given groupID
	err := repo.db.Where("group_id = ?", groupID).First(&courseGroup).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.CourseGroup{}, err
		}

		return models.CourseGroup{}, err
	}

	return courseGroup, nil
}
