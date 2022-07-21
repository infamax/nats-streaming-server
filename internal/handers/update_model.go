package handers

import (
	"github.com/gin-gonic/gin"
	"github.com/infamax/nats-streaming-server/internal/models"
	"net/http"
)

func (h *Handler) UpdateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		return
	}
	err := h.service.UpdateModelDB(&order)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "model doesn't exist",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, order)
}
