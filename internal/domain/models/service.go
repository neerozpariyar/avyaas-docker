package models

type Service struct {
	Timestamp

	Title       string `json:"title"`
	Description string `json:"description"`
}

/*
ServiceUrl is a model struct that represents the many2many relation between Service and Permission
Url models. The model is created separately(custom) as Url is a model dervied from roles and permissions
package.
*/
type ServiceUrl struct {
	ServiceID uint
	UrlID     uint
}
