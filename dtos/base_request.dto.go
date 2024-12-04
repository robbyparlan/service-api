package dtos

type BaseRequestDTO struct {
	Page     int    `json:"Page" validate:"gte=1"`     // Page number
	PageSize int    `json:"PageSize" validate:"gte=1"` // Number of items per page
	Search   string `json:"Search"`                    // Search term
}
