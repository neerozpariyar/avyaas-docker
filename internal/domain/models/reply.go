package models

type Reply struct {
	Timestamp

	DiscussionID uint       `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"discussionID"`
	Discussion   Discussion ` json:"-"`
	Reply        string     `json:"reply"`
	CourseID     uint       `json:"courseID"`
	Course       Course     `json:"-"`
	UserID       uint       `json:"userID"`
}

// func (m *Reply) AfterCreate(tx *gorm.DB) (err error) {
// 	return tx.Exec(`update discussions set reply_count =
// (select COUNT(id) from replies where replies.DiscussionID = Discussion.Id`, m.ID, m.ID).Error
// }
