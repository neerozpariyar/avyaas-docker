package presenter

type TeacherReferralList struct {
	Student      interface{} `json:"student"`      // student id, name
	Subscription interface{} `json:"subscription"` // package name, price, payment method
}

type TeacherReferralReq struct {
	Page     int
	PageSize int
}
