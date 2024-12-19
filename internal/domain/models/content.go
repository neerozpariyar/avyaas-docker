package models

type Content struct {
	Timestamp // TimeStamp represents the time when the content was created.

	Title       string `json:"title"` // Title is the title of the content.
	Description string `json:"description"`
	IsPremium   *bool  `json:"isPremium" gorm:"default:false"`
	ContentType string `json:"contentType"` // ContentType is the type of the content.
	Length      uint   `json:"length"`
	Level       string `json:"level"`
	Visibility  *bool  `json:"visibility" gorm:"default:false"`
	// CourseID    uint      `json:"courseID"`
	// Course      Course    `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	// ChapterID   uint      `json:"chapterID"`
	// Chapters    []Chapter `gorm:"many2many:chapter_contents, constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` //`gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"-"`
	Views     int    `json:"views" gorm:"default:0"`
	Url       string `json:"url"`
	Position  uint   `json:"-"`
	CreatedBy uint   `json:"-"` // who created this content ?
}

type ChapterContent struct { // for many2many relation only
	ChapterID uint `json:"chapterID" validate:"required"`
	ContentID uint `json:"contentID" validate:"required"`
	Position  uint `json:"position"`
}

// func (ChapterContent) BeforeCreate(db *gorm.DB) error { // for many2many relation only
// 	err := db.SetupJoinTable(&Chapter{}, "Contents", &ChapterContent{})
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
