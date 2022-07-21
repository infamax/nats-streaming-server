package handers

import (
	"github.com/gin-gonic/gin"
	"github.com/infamax/nats-streaming-server/internal/models"
	"net/http"
)

func (h *Handler) AddModelDB(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		return
	}
	err := h.service.AddModelDB(&order)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "model already exists in db",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, order)
}

func (h *Handler) AddModelCache(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		return
	}
	err := h.service.AddModelCache(&order)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "model already exist in db",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, order)

}
