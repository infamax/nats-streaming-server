package cache

import (
	"errors"
	"github.com/infamax/nats-streaming-server/internal/models"
)

func (c *Cache) Update(order *models.Order) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[order.OrderUid]
	if !ok {
		return errors.New("not found")
	}

	c.cache[order.OrderUid] = order
	return nil
}
