package cache

import "github.com/infamax/nats-streaming-server/internal/models"

func (c *Cache) Add(order *models.Order) {
	c.mu.Lock()
	c.cache[order.OrderUid] = order
	c.mu.Unlock()
}
