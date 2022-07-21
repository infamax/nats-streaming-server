package handers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/infamax/nats-streaming-server/internal/service"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) (*Handler, error) {
	if service == nil {
		return nil, errors.New("empty")
	}

	return &Handler{service: service}, nil
}

func (h *Handler) InitRoutes() *gin.Engine {
	h.refreshCache()
	router := gin.New()
	router.GET("/get_order_db/:id", h.GetByUUID)
	router.GET("/get_order_cache/:id", h.GetCache)
	router.PATCH("/update_order", h.UpdateOrder)
	router.POST("/create_order_db", h.AddModelDB)
	router.POST("/create_order_cache", h.AddModelCache)
	router.POST("/add_invalid_data", h.AddInvalidData)
	router.DELETE("/delete_order/:id", h.DeleteModel)
	return router
}

func (h *Handler) refreshCache() {
	orders, err := h.service.GetAllModels()
	if err != nil {
		return
	}

	for _, order := range orders {
		_ = h.service.AddModelCache(&order)
	}
}
