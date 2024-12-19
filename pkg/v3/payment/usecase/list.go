package usecase

import (
	"avyaas/internal/domain/presenter"

	"encoding/json"
)

func (u *usecase) ListPayments(request *presenter.PaymentListRequest) ([]presenter.PaymentListResponse, int, error) {
	payments, totalPage, err := u.repo.ListPayments(request)
	if err != nil {
		return nil, int(totalPage), err
	}

	var allPayments []presenter.PaymentListResponse
	for _, payment := range payments {
		var singlePayment presenter.PaymentListResponse

		// Convert the payment data to JSON format
		bData, err := json.Marshal(payment)
		if err != nil {
			return nil, 0, err
		}

		// Unmarshal the JSON data into the presenter.PaymentListResponse structure
		if err = json.Unmarshal(bData, &singlePayment); err != nil {
			return nil, 0, err
		}

		singlePayment.PaymentDate = payment.CreatedAt.UTC().Format("2006-01-02T15:04:05Z")

		user, err := u.accountRepo.GetUserByID(payment.UserID)
		if err != nil {
			return nil, 0, err
		}

		if user.RoleID != 4 {
			userData := make(map[string]interface{})
			userData["id"] = user.ID
			userData["name"] = user.FirstName + " " + user.LastName
			userData["phone"] = user.Phone
			singlePayment.User = userData
		}

		course, err := u.courseRepo.GetCourseByID(payment.CourseID)
		if err != nil {
			return nil, 0, err
		}

		courseData := make(map[string]interface{})
		courseData["id"] = course.ID
		courseData["courseID"] = course.CourseID
		singlePayment.Course = courseData

		pkg, err := u.packageRepo.GetPackageByID(payment.PackageID)
		if err != nil {
			return nil, 0, err
		}

		// packageData := make(map[string]interface{})
		// packageData["id"] = pkg.ID
		// packageData["tile"] = pkg.Title
		singlePayment.Package = pkg.Title

		allPayments = append(allPayments, singlePayment)
	}

	return allPayments, int(totalPage), nil
}
