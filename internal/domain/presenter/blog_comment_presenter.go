package presenter

type BlogCommentListReq struct {
	UserID   uint   `json:"userID"`
	BlogID   uint   `json:"blogID"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Search   string `json:"search"`
}

type BlogCommentListRes struct {
	Comment string      `json:"comment"`
	User    interface{} `json:"user"`
	Blog    interface{} `json:"blog"`
}
