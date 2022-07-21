package repository

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
)

type Repository interface {
	AddModel(ctx context.Context, order *models.Order) error
	GetByUUID(ctx context.Context, uuid string) (*models.Order, error)
	UpdateModel(ctx context.Context, order *models.Order) error
	DeleteModel(ctx context.Context, uuid string) error
	AddData(ctx context.Context, data string) (int, error)
	GetByID(ctx context.Context, id int) (string, error)
	UpdateData(ctx context.Context, id int, data string) error
	DeleteData(ctx context.Context, id int) error
}
