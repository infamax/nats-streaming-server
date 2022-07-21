package service

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
	"time"
)

func (s *service) GetModelDB(uuid string) (*models.Order, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.GetByUUID(ctx, uuid)
}

func (s *service) GetAllModels() ([]models.Order, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.GetAllModels(ctx)
}

func (s *service) GetModelCache(uuid string) (*models.Order, error) {
	return s.cache.Get(uuid)
}

func (s *service) GetData(id int) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.GetByID(ctx, id)
}
