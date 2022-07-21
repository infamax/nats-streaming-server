package service

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
	"time"
)

func (s *service) UpdateModelDB(order *models.Order) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.UpdateModel(ctx, order)
}

func (s *service) UpdateData(id int, data string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.repo.UpdateData(ctx, id, data)
}
