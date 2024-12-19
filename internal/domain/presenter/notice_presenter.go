package presenter

import "mime/multipart"

type NoticeCreateUpdatePresenter struct {
	ID          uint                  `json:"id"`
	CreatedBy   uint                  `json:"-"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	File        *multipart.FileHeader `json:"file"`
	CourseID    uint                  `json:"courseID"`
}
type NoticeListPresenter struct {
	ID          uint        `json:"id"`
	CreatedBy   uint        `json:"-"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	File        string      `json:"file"`
	Course      interface{} `json:"course"`
}

type NoticeListReq struct {
	Search   string `json:"search"`
	CourseID uint   `json:"courseID"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
