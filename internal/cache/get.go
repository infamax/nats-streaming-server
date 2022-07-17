package cache

import (
	"errors"
	"github.com/infamax/nats-streaming-server/internal/models"
)

func (c *Cache) Get(uuid string) (*models.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[uuid]
	if !ok {
		return nil, errors.New("not found")
	}
	return val, nil
}
