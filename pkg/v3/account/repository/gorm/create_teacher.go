package gorm

import (
	"avyaas/internal/core"
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"avyaas/utils/file"

	"log"
	"strings"

	"avyaas/utils"

	"gorm.io/gorm"
)

/*
CreateTeacher is a repository method responsible for creating a new teacher in the database.

Parameters:
  - user: An instance of the TeacherCreateUpdateRequest struct containing the necessary information
    for creating a new teacher.

Returns:
  - error: An error indicating any issues encountered during the teacher creation process.
    A nil error signifies a successful creation of the teacher.
*/
func (repo *Repository) CreateTeacher(data presenter.TeacherCreateUpdateRequest) error {
	var err error
	var user models.User

	transaction := repo.db.Begin()
	// Generate random password
	randPassword, err := utils.GenerateRandomPassword()
	if err != nil {
		return err
	}
	// Hash the user's password before storing it
	pwd, err := utils.HashPassword(randPassword)
	if err != nil {
		log.Fatal(err)
	}
	teacher := &models.User{
		FirstName:  data.FirstName,
		MiddleName: data.MiddleName,
		LastName:   data.LastName,
		Username:   data.Username,
		Gender:     models.Gender(data.Gender),
		Email:      data.Email,
		Phone:      data.Phone,
		RoleID:     3, // Teacher has role of ID 3.
		Verified:   true,
		Password:   pwd,
	}

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

		teacher.Image = urlObject
	}

	if err := transaction.Create(&teacher).Scan(&user).Error; err != nil {
		transaction.Rollback()
		return err
	}

	// Initialize the authority instance to access roles and permissions services
	auth := core.GetAuth(repo.db)

	// Assign the teacher role to the created user
	if err = auth.AssignRole(user.ID, 3); err != nil {
		transaction.Rollback()
		return err
	}
	for _, subjectID := range data.SubjectIDs {
		// Initiate the creation of teacher instance in the database
		err = repo.RegisterTeacher(user, subjectID, transaction)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}
	// Send email to the teacher user with email and password
	err = utils.TeacherAccountCreatedSMTP(user.Email, randPassword)
	if err != nil {
		transaction.Rollback()
		return err
	}
	if user.RoleID == 3 {
		if errList := auth.AssignUserPermissions(user.ID, []uint{3, 6, 8, 10, 14, 15, 21, 22, 26, 31, 36, 37, 45, 57, 58, 64, 65, 67, 68,
			73, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
			104, 140, 141, 142, 143}); len(errList) != 0 {
			log.Println(errList)
		}
	}
	// Commit the changes to the database if no error
	transaction.Commit()

	return nil
}

/*
RegisterTeacher is a repository method responsible for associating a teacher with a specific course and subject.
It is typically called within the context of a transaction initiated by the CreateTeacher method.

Parameters:
  - user: An instance of the User model representing the newly created teacher.
  - courseID: The identifier of the course associated with the teacher.
  - subjectID: The identifier of the subject associated with the teacher.
  - transaction: A reference to the database transaction to ensure atomicity of operations.

Returns:
  - error: An error indicating any issues encountered during the association of the teacher with a course and subject.
    A nil error signifies a successful association.
*/

func (repo *Repository) RegisterTeacher(user models.User, subjectID uint, transaction *gorm.DB) error {
	// Initiate the teacher creation

	// seed := rand.NewSource(time.Now().UnixNano())
	// rand.New(seed)

	// b := make([]byte, 6)
	// _, err := rand.Read(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// referral := base64.StdEncoding.EncodeToString(b)
	// fmt.Printf("referral: %v\n", referral)

	referral, _ := utils.GenerateReferralCode()

	// generateCode := rand.Intn(900000) + 100000
	// genCode := strconv.Itoa(generateCode)

	err := transaction.Create(&models.Teacher{
		Timestamp: models.Timestamp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},

		ReferralCode: strings.ToUpper(referral),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
