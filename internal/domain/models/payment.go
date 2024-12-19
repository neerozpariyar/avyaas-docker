package models

type Payment struct {
	Timestamp

	UserID             uint
	CourseID           uint
	PackageID          uint
	MerchantType       string
	TransactionID      string
	Discount           int
	Amount             int
	SubscriptionPeriod int
	Status             string
	TransactionUUID    string
}
