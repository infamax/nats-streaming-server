package service

import (
	"context"
	"time"
)

func (s *service) DeleteModelDB(uuid string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.DeleteModel(ctx, uuid)
}

func (s *service) DeleteData(id int) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.DeleteData(ctx, id)
}
