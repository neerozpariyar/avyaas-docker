package presenter

import "mime/multipart"

type BlogCreateUpdatePresenter struct {
	ID          uint                  `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Cover       *multipart.FileHeader `json:"cover"`
	CreatedBy   uint                  `json:"createdBy"`
	Tags        string                `json:"tags"`
	Views       uint                  `json:"views"`
	CourseID    uint                  `json:"courseID"`
	SubjectID   uint                  `json:"subjectID"`
}

type BlogListReq struct {
	CourseID  uint   `json:"courseID"`
	SubjectID uint   `json:"subjectID"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	Search    string `json:"search"`
}

type BlogListRes struct {
	ID      uint        `json:"id"`
	Title   string      `json:"title"`
	Cover   string      `json:"cover"`
	Tags    string      `json:"tags"`
	Course  interface{} `json:"course"`
	Subject interface{} `json:"subject"`
}

type BlogDetailsPresenter struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Tags        string `json:"tags"`
	Likes       uint   `json:"likes"`
	Description string `json:"description"`
	Views       uint   `json:"views"`
	Cover       string `json:"cover"`
	Comments    string `json:"comments"`
}
type BlogDetailsListRes struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
