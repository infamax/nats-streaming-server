package cache

import "errors"

func (c *Cache) Delete(uuid string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[uuid]
	if !ok {
		return errors.New("not found")
	}
	delete(c.cache, uuid)
	return nil
}
