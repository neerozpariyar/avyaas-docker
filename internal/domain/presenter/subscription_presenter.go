package presenter

type ListSubscriptionRequest struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	CourseID uint   `json:"courseID"`
	Search   string `json:"search"`
	UserID   uint   `json:"userID"`
}

type SubscriptionListResponse struct {
	ID            uint        `json:"id"`
	User          interface{} `json:"user"`
	Course        interface{} `json:"course"`
	Package       interface{} `json:"package"`
	PaymentID     uint        `json:"paymentID"`
	PaymentMethod string      `json:"paymentMethod"`
	TransactionID string      `json:"transactionID"`
	ReferralCode  string      `json:"referralCode"`
	ExpiryDate    string      `json:"expiryDate"`
}
