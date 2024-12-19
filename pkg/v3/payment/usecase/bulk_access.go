package usecase

import (
	"avyaas/internal/domain/models"
	"avyaas/internal/domain/presenter"
	"fmt"
	"math/rand"
	"time"

	"github.com/tealeg/xlsx"
)

func (uc *usecase) BulkAcessPayment(request *presenter.BulkAccessPaymentRequest) error {
	pkg, err := uc.packageRepo.GetPackageByID(request.PackageID)
	if err != nil {
		return err
	}

	src, err := request.File.Open()
	if err != nil {
		return err
	}

	xlsxFile, err := xlsx.OpenReaderAt(src, request.File.Size)
	if err != nil {
		return err
	}
	var phoneNumbers []string
	for _, sheet := range xlsxFile.Sheets {
		for i := 1; i < len(sheet.Rows); i++ {
			row := sheet.Rows[i]
			if len(row.Cells) >= 1 {
				cell := row.Cells[0] // assuming the phone number is in the 2nd column; as in 1st column is the index i.e. Phone
				phoneNumber := cell.String()
				phoneNumbers = append(phoneNumbers, phoneNumber)
			}
		}
	}
	var user *presenter.UserResponse

	for _, phone := range phoneNumbers {

		user, err = uc.accountRepo.GetUserByPhone(phone)
		if err != nil {
			return err
		}

		var randInvoice string

		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randInvoice = fmt.Sprintf("%09d", r.Intn(1000000000))
			existingPayment, _ := uc.repo.GetPaymentByTransactionID(randInvoice)

			if existingPayment == nil {
				break
			}
		}

		payment := models.Payment{
			UserID:             user.ID,
			PackageID:          request.PackageID,
			CourseID:           pkg.CourseID,
			Amount:             pkg.Price,
			Status:             "success",
			SubscriptionPeriod: pkg.Period,
			TransactionID:      randInvoice,
			MerchantType:       "bulk access",
		}

		subscription := presenter.SubscribePackageRequest{
			UserID:        user.ID,
			PackageID:     request.PackageID,
			TransactionID: randInvoice,
		}

		courseID, err := uc.packageRepo.GetCourseIDByPackageID(request.PackageID)
		if err != nil {
			return err
		}

		err = uc.courseRepo.EnrollInCourse(user.ID, courseID) //if student is not enrolled in the course, enroll them
		if err != nil {
			if err.Error() == "already enrolled" { //if student is already enrolled in the course, update the subscription status
				err = uc.packageRepo.SubscribePackage(subscription)
				if err != nil {
					return err
				}
			}
			return err
		}

		err = uc.packageRepo.SubscribePackage(subscription)
		if err != nil {
			return err
		}

		err = uc.repo.CreatePayment(&payment)
		if err != nil {
			return err
		}

	}
	return err
}
