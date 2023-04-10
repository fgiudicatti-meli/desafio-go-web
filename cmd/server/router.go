package server

import (
	"github.com/fgiudicatti-meli/desafio-go-web/cmd/handler"
	"github.com/fgiudicatti-meli/desafio-go-web/internal/domain"
	"github.com/fgiudicatti-meli/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func (router *Router) Setup(ticketList []domain.Ticket) {
	router.Engine.Use(gin.Logger())
	router.Engine.Use(gin.Recovery())

	router.SetTicketsRoutes(ticketList)
}

func (router *Router) SetTicketsRoutes(list []domain.Ticket) {
	// Setup components.
	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	controller := handler.NewTicketsHandler(service)

	// Set routes.
	group := router.Engine.Group("/tickets")
	{
		group.GET("getByCountry/:dest", controller.GetTicketsByCountry())
		group.GET("getAverage/:dest", controller.AverageDestination())
	}
}
