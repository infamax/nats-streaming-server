package handers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteModel(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteModelDB(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "model with such uuid not exist in db",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "model with successfully delete",
	})
}
