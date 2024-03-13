package services

import "cards-api/internal/helpers"

type cards struct{}

func NewCardsService() Cards {
	return &cards{}
}

func (s *cards) ValidateCard(input ValidateCardInput) error {
	err := helpers.ValidateCardNumber(input.Number)
	if err != nil {
		return err
	}

	err = helpers.ValidateCardExpirationDate(input.ExpirationMonth, input.ExpirationYear)
	if err != nil {
		return err
	}

	return nil
}
