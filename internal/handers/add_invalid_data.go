package handers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddInvalidData(c *gin.Context) {
	var data string
	if err := c.BindJSON(&data); err != nil {
		return
	}
	_, err := h.service.AddData(data)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "model already exist",
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, data)
}
