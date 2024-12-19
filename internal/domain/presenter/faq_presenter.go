package presenter

type FAQListReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
