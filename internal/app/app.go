package app

import (
	"cards-api/internal/api"
	"cards-api/internal/server"
	"cards-api/internal/services"
)

func Run() {
	router := server.NewServer()

	group := router.Group("api")

	cardsService := services.NewCardsService()

	api.UseCards(group, cardsService)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
