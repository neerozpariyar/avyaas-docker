package presenter

import "mime/multipart"

type InitateKhaltiPaymentRequest struct {
	UserID            uint   `json:"userID"`
	ReturnUrl         string `json:"returnUrl" validate:"required"`
	WebsiteUrl        string `json:"websiteUrl" validate:"required"`
	Amount            int    `json:"amount" validate:"required"`
	PurchaseOrderID   string `json:"purchaseOrderID" validate:"required"`
	PurchaseOrderName string `json:"purchaseOrderName" validate:"required"`
}

type InitiateKhaltiPaymentResponse struct {
	PIDx       string `json:"pidx"`
	PaymentUrl string `json:"payment_url"`
	ExpiresAt  string `json:"expires_at"`
	ExpiresIn  int    `json:"expires_in"`
}

// VerifyPaymentRequest is a presenter struct for the request payload of the verify payment endpoint.
type VerifyPaymentRequest struct {
	UserID          uint   `json:"userID"`
	CourseID        uint   `json:"courseID" validate:"required"`
	PackageID       uint   `json:"packageID" validate:"required"`
	PaymentMethod   string `json:"paymentMethod" validate:"required"`
	Amount          int    `json:"amount" validate:"required"`
	TransactionUUID string `json:"transactionUuid" validate:"required"` // transaction_uuid for eSewa and pidx for Khalti
	TransactionID   string `json:"transactionID" validate:"required"`
}

// VerifyPaymentStatus is a presenter struct to map the payment verify response returned from eSewa.
type VerifyPaymentStatus struct {
	ProductCode     string  `json:"product_code"`     // eSewa merchant secret key
	TransactionUUID string  `json:"transaction_uuid"` // unique code to identify the transaction
	TotalAmount     float64 `json:"total_amount"`     // total amount of transaction
	Status          string  `json:"status"`           // status of the transaction
	ReferenceID     string  `json:"ref_id"`           // transaction/receipt id

	// Khalti specific response
	PIDx          string  `json:"pidx"`           // Khalti payment ID
	TransactionID string  `json:"transaction_id"` // Khalti transaction ID
	Fee           float32 `json:"fee"`            // Khalti transaction fee
	Refunded      bool    `json:"refunded"`       // Khalti's refund status of the payment
}

type PaymentListRequest struct {
	UserID   uint
	Page     int
	CourseID int
	Search   string
	PageSize int
}

type PaymentListResponse struct {
	ID                 uint        `json:"id"`
	PaymentDate        string      `json:"paymentDate"`
	User               interface{} `json:"user"`
	Course             interface{} `json:"course"`
	Package            string      `json:"package"`
	MerchantType       string      `json:"merchantType"`
	TransactionID      string      `json:"transactionID"`
	Discount           int         `json:"discount"`
	Amount             int         `json:"amount"`
	SubscriptionPeriod int         `json:"subscriptionPeriod"`
	Status             string      `json:"status"`
	TransactionUUID    string      `json:"transactionUuid"`
}

type ManualPaymentRequest struct {
	UserID        uint   `json:"userID"`
	PackageID     uint   `json:"packageID" validate:"required"`
	InvoiceNumber string `json:"invoiceNumber" validate:"required"`
	// PaymentMethod string `json:"paymentMethod" validate:"required"`
}

type BulkAccessPaymentRequest struct {
	PackageID uint                  `json:"packageID" validate:"required"`
	File      *multipart.FileHeader `json:"file" validate:"required"`
}
