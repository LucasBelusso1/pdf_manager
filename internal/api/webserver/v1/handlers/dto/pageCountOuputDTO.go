package dto

type PageCountOutputDTO struct {
	Name         string `json:"name"`
	PagesQty     int    `json:"pagesQty"`
	Error        bool   `json:"error"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
