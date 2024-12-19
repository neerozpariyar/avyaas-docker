package gorm

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"

	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func (repo *Repository) CreateSubscription(transaction *gorm.DB, data models.Subscription) error {
	return transaction.Create(&data).Error
}

func (repo *Repository) CheckUserSubscription(data models.Subscription) error {
	return repo.db.Where("user_id = ? AND course_id = ? AND package_id = ?", data.UserID, data.CourseID, data.PackageID).First(&models.Subscription{}).Error
}

func (repo *Repository) SubscribePackage(request presenter.SubscribePackageRequest) error {
	var err error
	transaction := repo.db.Begin()

	pkg, err := repo.GetPackageByID(request.PackageID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	serviceIDs, err := repo.packageTypeRepo.GetPackageTypeServices(pkg.PackageTypeID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	expiryDate := time.Now().AddDate(0, 0, pkg.Period)
	request.ExpiryDate = &expiryDate

	subscriptionRequest := models.Subscription{
		UserID:        request.UserID,
		CourseID:      pkg.CourseID,
		PackageID:     request.PackageID,
		PaymentID:     request.PaymentID,
		PaymentMethod: request.PaymentMethod,
		TransactionID: request.TransactionID,
		ExpiryDate:    &expiryDate,
	}

	if err = repo.CheckUserSubscription(subscriptionRequest); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = repo.CreateSubscription(transaction, subscriptionRequest); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		return fmt.Errorf("already subscribed to the course packageid : %v", request.PackageID)
	}

	err = transaction.Where("user_id = ? AND package_id = ?", request.UserID, request.PackageID).Delete(&models.ReferralInTransaction{}).Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	for idx := range serviceIDs {
		switch serviceIDs[idx] {
		case 1:
			println("course")
			err = repo.MakeCourseSubscription(request.UserID, pkg.CourseID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		case 2:
			println("test series")
			err = repo.MakeTestSeriesSubscription(request.UserID, pkg.TestSeriesID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		case 3:
			println("test")
			err = repo.MakeSingleTestSubscription(request.UserID, pkg.CourseID, pkg.TestID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		case 4:
			println("live group")
			err = repo.MakeLiveGroupSubscription(request.UserID, pkg.CourseID, pkg.LiveGroupID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		case 5:
			println("live")
			err = repo.MakeSingleLiveSubscription(request.UserID, pkg.CourseID, pkg.LiveID, transaction)
			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	transaction.Commit()

	return err
}

func (repo *Repository) MakeCourseSubscription(userID, courseID uint, transaction *gorm.DB) error {
	studentCourse, err := repo.courseRepo.CheckStudentCourse(userID, courseID)
	if err != nil {
		return err
	}

	err = transaction.Debug().Model(&models.StudentCourse{}).Where("id = ?", studentCourse.ID).Update("paid", true).Error
	if err != nil {
		return err
	}

	err = transaction.Debug().Model(&models.StudentContent{}).Where("course_id = ?", studentCourse.CourseID).Update("paid", true).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) MakeTestSeriesSubscription(userID, testSeriesID uint, transaction *gorm.DB) error {
	err := transaction.Create(&models.StudentTestSeries{
		UserID:       userID,
		TestSeriesID: testSeriesID,
	}).Error

	if err != nil {
		return err
	}

	var tests []models.Test
	err = repo.db.Where("test_series_id = ?", testSeriesID).Find(&tests).Error
	if err != nil {
		return err
	}

	for _, eachTest := range tests {
		err = transaction.Debug().Create(&models.StudentTest{
			UserID:   userID,
			TestID:   eachTest.ID,
			CourseID: eachTest.CourseID,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) MakeSingleTestSubscription(userID, courseID, testID uint, transaction *gorm.DB) error {
	err := transaction.Create(&models.StudentTest{
		UserID:   userID,
		TestID:   testID,
		CourseID: courseID,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) MakeSingleLiveSubscription(userID, courseID, liveID uint, transaction *gorm.DB) error {
	var live models.Live
	err := repo.db.Where("id = ?", liveID).First(&live).Error
	if err != nil {
		return err
	}

	err = transaction.Create(&models.StudentLive{
		UserID:    userID,
		CourseID:  courseID,
		SubjectID: live.SubjectID,
		LiveID:    liveID,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) MakeLiveGroupSubscription(userID, courseID, liveGroupID uint, transaction *gorm.DB) error {
	err := transaction.Create(&models.StudentLiveGroup{
		UserID:      userID,
		CourseID:    courseID,
		LiveGroupID: liveGroupID,
	}).Error

	if err != nil {
		return err
	}

	var lives []models.Live
	err = repo.db.Where("live_group_id = ?", liveGroupID).Find(&lives).Error
	if err != nil {
		return err
	}

	for _, eachLive := range lives {
		err = transaction.Debug().Create(&models.StudentLive{
			UserID:    userID,
			CourseID:  eachLive.CourseID,
			SubjectID: eachLive.SubjectID,
			LiveID:    eachLive.ID,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

// func (repo *repository) SubscribePackage(request presenter.SubscribePackageRequest) error {
// 	var err error
// 	transaction := repo.db.Begin()

// 	pkg, err := repo.GetPackageByID(request.PackageID)
// 	if err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	expiryDate := time.Now().AddDate(0, 0, pkg.Period)
// 	request.ExpiryDate = &expiryDate

// 	subscriptionRequest := models.Subscription{
// 		UserID:        request.UserID,
// 		CourseID:      pkg.CourseID,
// 		PackageID:     request.PackageID,
// 		PaymentID:     request.PaymentID,
// 		PaymentMethod: request.PaymentMethod,
// 		TransactionID: request.TransactionID,
// 		ExpiryDate:    &expiryDate,
// 	}

// 	if err = repo.CheckUserSubscription(subscriptionRequest); err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			if err = repo.CreateSubscription(transaction, subscriptionRequest); err != nil {
// 				return err
// 			}
// 		} else {
// 			return err
// 		}
// 	} else {
// 		return fmt.Errorf("already subscribed to the course packageid : %v", request.PackageID)
// 	}

// 	err = transaction.Where("user_id = ? AND package_id = ?", request.UserID, request.PackageID).Delete(&models.ReferralInTransaction{}).Error
// 	if err != nil {
// 		transaction.Rollback()
// 		return err
// 	}

// 	switch pkg.PackageTypeID {
// 	case 1:
// 		err = repo.MakeCourseSubscription(request.UserID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		err = repo.MakeTestSeriesSubscription(request.UserID, pkg.TestSeriesID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		err = repo.MakeLiveGroupSubscription(request.UserID, pkg.CourseID, pkg.LiveGroupID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 2:
// 		err = repo.MakeCourseSubscription(request.UserID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		err = repo.MakeTestSeriesSubscription(request.UserID, pkg.TestSeriesID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 3:
// 		err = repo.MakeCourseSubscription(request.UserID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		err = repo.MakeLiveGroupSubscription(request.UserID, pkg.LiveGroupID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 4:
// 		err = repo.MakeTestSeriesSubscription(request.UserID, pkg.TestSeriesID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}

// 		err = repo.MakeLiveGroupSubscription(request.UserID, pkg.CourseID, pkg.LiveGroupID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 5:
// 		err = repo.MakeCourseSubscription(request.UserID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 6:
// 		err = repo.MakeTestSeriesSubscription(request.UserID, pkg.TestSeriesID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 7:
// 		err = repo.MakeLiveGroupSubscription(request.UserID, pkg.LiveGroupID, pkg.CourseID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 8:
// 		err = repo.MakeSingleTestSubscription(request.UserID, pkg.CourseID, pkg.TestID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	case 9:
// 		err = repo.MakeSingleLiveSubscription(request.UserID, pkg.CourseID, pkg.LiveID, transaction)
// 		if err != nil {
// 			transaction.Rollback()
// 			return err
// 		}
// 	default:
// 		return errors.New("invalid package type")
// 	}

// 	transaction.Commit()

// 	return err

// 	// var allUrlIDs []uint
// 	// for _, service := range packageData.Services {
// 	// 	urlIDs, err := repo.serviceRepo.GetUrlIDsByServiceID(service.ID)
// 	// 	if err != nil {
// 	// 		transaction.Rollback()
// 	// 		return err
// 	// 	}

// 	// 	allUrlIDs = append(allUrlIDs, urlIDs...)
// 	// }

// 	// var urlPaths []string
// 	// for _, urlId := range allUrlIDs {
// 	// 	var path string
// 	// 	err = repo.db.Select("path").Table("urls").Where("id = ?", urlId).Scan(&path).Error
// 	// 	if err != nil {
// 	// 		transaction.Rollback()
// 	// 		return err
// 	// 	}

// 	// 	urlPaths = append(urlPaths, path)
// 	// }

// 	// if utils.Contains(urlPaths, "/course/list/") {
// 	// 	studentCourse, err := repo.CheckStudentCourse(request.UserID, request.CourseID)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}

// 	// 	err = transaction.Debug().Model(&models.StudentCourse{}).Where("id = ?", studentCourse.ID).Update("paid", true).Error
// 	// 	if err != nil {
// 	// 		transaction.Rollback()
// 	// 		return err
// 	// 	}

// 	// 	err = transaction.Debug().Model(&models.StudentContent{}).Where("course_id = ?", studentCourse.CourseID).Update("paid", true).Error
// 	// 	if err != nil {
// 	// 		transaction.Rollback()
// 	// 		return err
// 	// 	}

// 	// return err
// 	// if err = repo.EnrollInCourse(request, transaction); err != nil {
// 	// 	transaction.Rollback()
// 	// 	return err
// 	// }
// 	// }

// 	// Assign url permissions to the user
// 	// if errList := auth.AssignUserPermissions(request.UserID, allUrlIDs); len(errList) != 0 {
// 	// 	// panic(errList)
// 	// 	println(".................................")
// 	// }

// 	// transaction.Commit()

// 	// return err
// }
