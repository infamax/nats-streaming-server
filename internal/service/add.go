package service

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
	"time"
)

func (s *service) AddModel(order *models.Order) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.AddModel(ctx, order)
}

func (s *service) AddData(data string) (int, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.AddData(ctx, data)
}
