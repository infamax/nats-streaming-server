package cache

import (
	"github.com/infamax/nats-streaming-server/internal/models"
	"sync"
)

type Cache struct {
	mu    *sync.RWMutex
	cache map[string]*models.Order
}

func New() *Cache {
	return &Cache{
		mu:    &sync.RWMutex{},
		cache: map[string]*models.Order{},
	}
}
