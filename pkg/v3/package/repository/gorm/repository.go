package gorm

import (
	"avyaas/internal/domain/interfaces"

	"gorm.io/gorm"
)

/*
repository represents the package repository, which encapsulates the Gorm database connection for
handling package related data operations.
*/
type Repository struct {
	db              *gorm.DB
	courseRepo      interfaces.CourseRepository
	packageTypeRepo interfaces.PackageTypeRepository
}

/*
New creates and returns a new instance of the package repository. It takes a Gorm database connection
as a parameter, which is used for data access.
*/
func New(db *gorm.DB, courseRepo interfaces.CourseRepository, packageTypeRepo interfaces.PackageTypeRepository) *Repository {
	return &Repository{
		db:              db,
		courseRepo:      courseRepo,
		packageTypeRepo: packageTypeRepo,
	}
}
