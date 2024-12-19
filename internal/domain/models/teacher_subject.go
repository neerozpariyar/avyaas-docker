package models

type TeacherSubject struct {
	UserID    uint `json:"userID"`
	SubjectID uint `json:"subjectID"`
	Subject   Subject
	CourseID  uint `json:"unitID"`
}

// type Teacher struct {
// 	gorm.Model
// 	Name     string
// 	Subjects []Subject `gorm:"many2many:teacher_subjects;"`
// }

// type Subject struct {
// 	gorm.Model
// 	Name     string
// 	Teachers []Teacher `gorm:"many2many:teacher_subjects;"`
// }
