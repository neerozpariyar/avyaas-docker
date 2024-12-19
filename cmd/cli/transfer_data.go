package main

import (
	"avyaas/internal/domain/models"
	"avyaas/utils"

	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var oldModuleIDs []string
var oldSubjectIDs []sql.NullInt32
var oldUnitIDs []int
var oldChapterIDs []int

func main() {
	oldDSN := "root:root@tcp(localhost:3306)/ebidhyaJan3?charset=utf8mb4&parseTime=True&loc=Local"
	oldDB, err := gorm.Open(mysql.Open(oldDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	newDSN := "root:root@tcp(localhost:3306)/avyaas_go?charset=utf8mb4&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(newDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := migrateCourseGroup(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateCourse(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateSubject(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateUnit(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateChapter(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateContent(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	if err := migrateUser(oldDB, newDB); err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+][+] Data migration completed successfully! [+][+]")
	fmt.Printf("oldModuleIDs: %v\n", oldModuleIDs)
	fmt.Printf("oldSubjectIDs: %v\n", oldSubjectIDs)
	fmt.Printf("oldUnitIDs: %v\n", oldUnitIDs)
	fmt.Printf("oldChapterIDs: %v\n", oldChapterIDs)
}

func migrateCourseGroup(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO course_groups (id, created_at, updated_at, title, group_id, description)
		VALUES (?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, name, groupId, description FROM courseGroup").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from courseGroup: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId int
		var oldName, oldGroupID, oldDescription string
		createdAt := time.Now()

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldName, &oldGroupID, &oldDescription); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		// Execute the custom INSERT query with mapped values
		err := newDB.Debug().Exec(insertQuery,
			uint(oldId),
			createdAt,
			createdAt,
			oldName,
			oldGroupID,
			oldDescription).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into course_group: %v", err)
		}
	}

	return nil
}

func migrateCourse(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO courses (id, created_at, updated_at, course_id, title, description, available, course_group_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, module_id, name, description, available, courseGroupId FROM course").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from course: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId, oldCourseGroupId int
		var oldModuleId, oldName, oldDescription string
		var oldAvailable bool
		createdAt := time.Now()

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldModuleId, &oldName, &oldDescription, &oldAvailable, &oldCourseGroupId); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		// Execute the custom INSERT query with mapped values
		err := newDB.Debug().Exec(insertQuery,
			uint(oldId),
			createdAt,
			createdAt,
			oldModuleId,
			oldName,
			oldDescription,
			oldAvailable,
			uint(oldCourseGroupId)).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into course: %v", err)
		}
	}

	return nil
}

func migrateSubject(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO subjects (id, created_at, updated_at, subject_id, title, description, course_id)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, subject_name, module_id FROM subject").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from subject: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId int
		var oldName, oldModuleId string
		createdAt := time.Now()

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldName, &oldModuleId); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		var course models.Course
		err := newDB.Model(&models.Course{}).Where("course_id = ?", oldModuleId).First(&course).Error
		if err != nil {
			if !utils.Contains(oldModuleIDs, oldModuleId) {
				oldModuleIDs = append(oldModuleIDs, oldModuleId)
			}
			continue
			// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
		}

		// Execute the custom INSERT query with mapped values
		err = newDB.Debug().Exec(insertQuery, uint(oldId), createdAt, createdAt, "", oldName, "", course.ID).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into course: %v", err)
		}
	}

	return nil
}

func migrateUnit(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO units (id, created_at, updated_at, title, description, subject_id)
		VALUES (?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, unit_name, subject_id FROM unit").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from subject: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId int
		var oldSubjectId sql.NullInt32
		var oldName string
		createdAt := time.Now()

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldName, &oldSubjectId); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		if oldSubjectId.Valid {
			var subject models.Subject
			err := newDB.Model(&models.Subject{}).Where("id = ?", oldSubjectId).First(&subject).Error
			if err != nil {
				if !utils.ContainsNullSQL(oldSubjectIDs, oldSubjectId) {
					oldSubjectIDs = append(oldSubjectIDs, oldSubjectId)
				}
				continue
				// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
			}

			// Execute the custom INSERT query with mapped values
			err = newDB.Debug().Exec(insertQuery, uint(oldId), createdAt, createdAt, oldName, "", subject.ID).Error
			if err != nil {
				return fmt.Errorf("failed to insert data into unit: %v", err)
			}
		} else {
			continue
		}
	}

	return nil
}

func migrateChapter(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO chapters (id, created_at, updated_at, title, unit_id)
		VALUES (?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, chapter_name, unit_id FROM chapters").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from subject: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId, oldUnitId int
		var oldName string
		createdAt := time.Now()

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldName, &oldUnitId); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		var unit models.Unit
		err := newDB.Model(&models.Unit{}).Where("id = ?", oldUnitId).First(&unit).Error
		if err != nil {
			if !utils.ContainsInt(oldUnitIDs, oldUnitId) {
				oldUnitIDs = append(oldUnitIDs, oldUnitId)
			}
			continue
			// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
		}

		// Execute the custom INSERT query with mapped values
		err = newDB.Debug().Exec(insertQuery, uint(oldId), createdAt, createdAt, oldName, unit.ID).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into unit: %v", err)
		}
	}

	return nil
}

func migrateContent(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO contents (
			id,
			created_at,
			updated_at,
			title,
			description,
			is_premium,
			content_type,
			length,
			level,
			visibility,
			views,
			url,
			created_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, ent_date, title, description, premium, length, level, module_id, chapterId, view, link, userId FROM video").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from video: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldEntDate time.Time // created date
		var oldId, oldChapterId, oldLength, oldViews, oldCreatedBy int
		var oldTitle, oldDescription, oldLevel, oldUrl, oldCourseId string
		var oldIsPremium bool

		// Scan values from old_table
		if err := rows.Scan(&oldId,
			&oldEntDate,
			&oldTitle,
			&oldDescription,
			&oldIsPremium,
			&oldLength,
			&oldLevel,
			&oldCourseId,
			&oldChapterId,
			&oldViews,
			&oldUrl,
			&oldCreatedBy); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		var course models.Course
		err = newDB.Model(&models.Course{}).Where("course_id = ?", oldCourseId).First(&course).Error
		if err != nil {
			continue
			// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
		}

		var chapter *models.Chapter
		err = newDB.Model(&models.Chapter{}).Where("id = ?", oldChapterId).First(&chapter).Error
		if err != nil {
			if !utils.ContainsInt(oldChapterIDs, oldChapterId) {
				oldChapterIDs = append(oldChapterIDs, oldChapterId)
			}
			continue
			// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
		}

		// Execute the custom INSERT query with mapped values
		err := newDB.Debug().Exec(insertQuery,
			uint(oldId),
			oldEntDate,
			oldEntDate,
			oldTitle,
			oldDescription,
			oldIsPremium,
			"VIDEO",
			oldLength,
			oldLevel,
			true,
			// course.ID,
			// chapter.ID,
			oldViews,
			oldUrl,
			oldCreatedBy,
		).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into unit: %v", err)
		}

		if chapter != nil {
			err = newDB.Create(&models.ChapterContent{
				ChapterID: uint(chapter.ID),
				ContentID: uint(oldId),
				Position:  0,
			}).Error

			if err != nil {
				return fmt.Errorf("failed to insert content: '%v' into chapter: %v", oldId, chapter.ID)
			}
		}
	}

	return nil
}

func migrateUser(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertQuery := `
		INSERT INTO users (id,
							created_at,
							updated_at,
							first_name,
							middle_name,
							last_name,
							username,
							gender,
							email,
							phone,
							role_id,
							verified,
							image,
							password,
							oauth_id,
							facebook_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT id, created_at, updated_at, email, phone, role, verified, password FROM user").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from courseGroup: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldId int
		var oldCreatedAt, oldUpdatedAt time.Time
		var oldRole, oldPassword string
		var oldEmail, oldPhone sql.NullString
		var oldVerified bool

		// Scan values from old_table
		if err := rows.Scan(&oldId, &oldCreatedAt, &oldUpdatedAt, &oldEmail, &oldPhone, &oldRole, &oldVerified, &oldPassword); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		var username string
		if oldEmail.Valid {
			username = (strings.Split(oldEmail.String, "@"))[0]
		} else {
			username = oldPhone.String
		}

		var roleId uint
		if oldRole == "administration" {
			roleId = 1
		} else if oldRole == "student" {
			roleId = 4
		}

		var newEmail string
		if oldEmail.Valid {
			newEmail = oldEmail.String
		}

		var newPhone string
		if oldPhone.Valid {
			newPhone = oldPhone.String
		}

		// Execute the custom INSERT query with mapped values
		err := newDB.Debug().Exec(insertQuery,
			uint(oldId),
			oldCreatedAt,
			oldUpdatedAt,
			"",
			"",
			"",
			username,
			"",
			newEmail,
			newPhone,
			roleId,
			oldVerified,
			"",
			oldPassword,
			"",
			"").Error
		if err != nil {
			return fmt.Errorf("failed to insert data into user: %v", err)
		}

		if oldRole == "administration" {
			err := newDB.Exec("INSERT INTO user_role_ints (user_id, role_id) VALUES (?, 1)", oldId).Error
			if err != nil {
				return fmt.Errorf("failed to assign admin role to the userID: %v", oldId)
			}

			err = newDB.Exec("INSERT INTO user_permission_ints (user_id, permission_id) VALUES (?, 1)", oldId).Error
			if err != nil {
				return fmt.Errorf("failed to assign admin permission to the userID: %v", oldId)
			}
		} else if oldRole == "student" {
			err := newDB.Exec("INSERT INTO user_role_ints (user_id, role_id) VALUES (?, 4)", oldId).Error
			if err != nil {
				return fmt.Errorf("failed to assign student role to the userID: %v", oldId)
			}
		}
	}

	return nil
}

func migrateTeacher(oldDB, newDB *gorm.DB) error {
	// Define your custom insert query mapping old fields to new fields
	insertUserQuery := `
		INSERT INTO users (created_at,
						updated_at,
						first_name,
						middle_name,
						last_name,
						username,
						gender,
						email,
						phone,
						role_id,
						verified,
						image,
						password,
						oauth_id,
						facebook_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	insertTeacherQuery := `
			INSERT INTO teachers (
				id,
				created_at,
				updated_at,
				course_id,
				subject_id
			)
			VALUES (?, ?, ?, ?, ?);
	`

	// Fetch all data from old_table
	rows, err := oldDB.Raw("SELECT created_at, updated_at, first_name, middle_name, last_name, gender, email, phone, subject_id, module_id FROM teacher").Rows()
	if err != nil {
		return fmt.Errorf("failed to fetch data from teacher: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var oldSubjectId int
		var oldCreatedAt, oldUpdatedAt time.Time
		var oldFirstName, oldMiddleName, oldLastName, oldGender, oldEmail, oldPhone, oldModuleId string // module_id = course_id

		// Scan values from old_table
		if err := rows.Scan(&oldCreatedAt, &oldUpdatedAt, &oldFirstName, &oldMiddleName, &oldLastName, &oldGender, &oldEmail, &oldPhone, &oldSubjectId, &oldModuleId); err != nil {
			return fmt.Errorf("failed to scan row data: %v", err)
		}

		var course models.Course
		err := newDB.Model(&models.Course{}).Where("course_id = ?", oldModuleId).First(&course).Error
		if err != nil {
			continue
			// return fmt.Errorf("could not find corresponding course for module ID %s: %w", oldModuleId, err)
		}

		var newUser models.User
		// Execute the custom INSERT query with mapped values
		err = newDB.Debug().Exec(insertUserQuery,
			oldCreatedAt,
			oldUpdatedAt,
			oldFirstName,
			oldMiddleName,
			oldLastName,
			"",
			oldGender,
			oldEmail,
			oldPhone,
			3,
			false,
			"",
			"",
			"",
			"").Scan(&newUser).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into user: %v", err)
		}

		fmt.Printf("newUser: %v\n", newUser)

		// Execute the custom INSERT query with mapped values
		err = newDB.Debug().Exec(insertTeacherQuery,
			newUser.ID,
			oldCreatedAt,
			oldUpdatedAt,
			course.ID,
			oldSubjectId,
		).Error
		if err != nil {
			return fmt.Errorf("failed to insert data into teacher: %v", err)
		}

		err = newDB.Exec("INSERT INTO user_role_ints (user_id, role_id) VALUES (?, 3)", newUser.ID).Error
		if err != nil {
			return fmt.Errorf("failed to assign teacher role to the userID: %v", newUser.ID)
		}
	}

	return nil
}
