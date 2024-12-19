package gorm

import (
	"avyaas/internal/domain/models"
	"log"
)

func (repo *Repository) CreateLive(data models.Live) error {
	transaction := repo.db.Begin()

	// if data.Type == 2 {
	// 	endDateTime := data.StartTime.Add(time.Minute * time.Duration(data.Duration))
	// 	data.EndDateTime = &endDateTime
	// }

	var newLive *models.Live
	err := transaction.Create(&models.Live{
		Topic:       data.Topic,
		LiveGroupID: data.LiveGroupID,
		CourseID:    data.CourseID,
		SubjectID:   data.SubjectID,
		StartTime:   data.StartTime,
		EndDateTime: data.EndDateTime,
		Duration:    data.Duration,
		IsLive:      data.IsLive,
		Type:        data.Type,
		MeetingID:   data.MeetingID,
		MeetingPwd:  data.MeetingPwd,
		Email:       data.Email,
		IsFree:      data.IsFree,
		IsPackage:   data.IsPackage,
	}).Scan(&newLive).Error
	if err != nil {
		transaction.Rollback()
		log.Printf("Error creating local database record: %+v\n", err)
		return err
	}

	if *data.IsPackage {
		err = transaction.Create(&models.Package{
			Title:         data.Topic,
			PackageTypeID: 9,
			CourseID:      data.CourseID,
			LiveID:        newLive.ID,
			LiveGroupID:   data.LiveGroupID,
			Price:         data.Price,
		}).Error
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	transaction.Commit()
	return nil
}
