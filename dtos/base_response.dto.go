package dtos

type BaseResponseDTO[T any] struct {
	Values      []T   `json:"Values"`
	CurrentPage int64 `json:"CurrentPage"`
	PageSize    int64 `json:"PageSize"`
	TotalItems  int64 `json:"TotalItems"`
	TotalPages  int64 `json:"TotalPages"`
}
