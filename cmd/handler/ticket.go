package handler

import (
	"net/http"

	"github.com/fgiudicatti-meli/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type TicketsHandler struct {
	service tickets.Service
}

func NewTicketsHandler(s tickets.Service) *TicketsHandler {
	return &TicketsHandler{
		service: s,
	}
}

func (h *TicketsHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		ticketsByDest, err := h.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, ticketsByDest)
	}
}

func (h *TicketsHandler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := h.service.GetAverageByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
