package models

/*
User represents a database model for storing user-related information, extending the Timestamp
model.
*/
type User struct {
	Timestamp

	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	Gender      Gender `json:"gender"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	RoleID      int    `json:"-"`
	Verified    bool   `gorm:"default:false" json:"-"`
	Image       string `json:"image"`
	Password    string `json:"password"`
	CollegeName string `json:"collegeName"`
	OauthID     string `json:"-"`
	FacebookID  string `json:"-"`
	// Courses    []Course `gorm:"many2many:user_courses;" json:"-"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
	Other  Gender = "Others"
)
