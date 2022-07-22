package cache

import (
	"errors"
	"github.com/infamax/nats-streaming-server/internal/models"
)

func (c *Cache) Add(order *models.Order) error {
	defer c.mu.Unlock()
	c.mu.Lock()
	_, ok := c.cache[order.OrderUid]
	if ok {
		return errors.New("model already exists in cache")

	}
	c.cache[order.OrderUid] = order
	return nil
}
