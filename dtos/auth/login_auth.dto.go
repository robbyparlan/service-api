package dtos

type LoginAuthDTO struct {
	Username string `json:"Username" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type LoginAuthResponseDTO[T any] struct {
	User        T      `json:"User"`
	AccessToken string `json:"AccessToken"`
}
