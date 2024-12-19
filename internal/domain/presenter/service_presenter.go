package presenter

type ServiceCreateUpdateRequest struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	// UrlIDs      []uint `json:"urlIDs" validate:"required"`
}
