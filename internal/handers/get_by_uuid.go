package handers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

func (h *Handler) GetByUUID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetModelDB(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"object": "not found",
		})
		return
	}
	_ = tmpl.Execute(c.Writer, order)
	//c.JSON(http.StatusOK, order)
}

func (h *Handler) GetCache(c *gin.Context) {
	id := c.Param("id")
	order, err := h.service.GetModelCache(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"object": "not found",
		})
		return
	}
	c.JSON(http.StatusOK, order)

}
