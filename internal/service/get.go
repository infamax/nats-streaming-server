package service

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
	"time"
)

func (s *service) GetModel(uuid string) (*models.Order, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.GetByUUID(ctx, uuid)
}

func (s *service) GetData(id int) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.GetByID(ctx, id)
}
