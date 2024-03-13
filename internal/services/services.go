package services

type ValidateCardInput struct {
	Number          string `json:"number" binding:"required"`
	ExpirationMonth string `json:"expirationMonth" binding:"required"`
	ExpirationYear  string `json:"expirationYear" binding:"required"`
}

type Cards interface {
	ValidateCard(input ValidateCardInput) error
}
