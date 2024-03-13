package api

import (
	"cards-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UseCards(group *gin.RouterGroup, cardsService services.Cards) {
	cards := group.Group("cards")

	cards.POST("validate", validateCard(cardsService))
}

type ValidateCardResponse struct {
	Valid bool      `json:"valid"`
	Error *ApiError `json:"error,omitempty"`
}

func validateCard(cardsService services.Cards) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input services.ValidateCardInput
		if err := ctx.BindJSON(&input); err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		var response ValidateCardResponse

		if err := cardsService.ValidateCard(input); err != nil {
			response.Error = &ApiError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		} else {
			response.Valid = true
		}

		ctx.JSON(http.StatusOK, response)
	}
}
