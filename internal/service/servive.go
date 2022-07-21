package service

import (
	"errors"
	"github.com/infamax/nats-streaming-server/internal/cache"
	"github.com/infamax/nats-streaming-server/internal/models"
	"github.com/infamax/nats-streaming-server/internal/repository"
)

type Service interface {
	AddModel(order *models.Order) error
	GetModel(uuid string) (*models.Order, error)
	UpdateModel(order *models.Order) error
	DeleteModel(uuid string) error
	AddData(data string) (int, error)
	GetData(id int) (string, error)
	UpdateData(id int, data string) error
	DeleteData(id int) error
}

type service struct {
	repo  repository.Repository
	cache *cache.Cache
}

func New(repo repository.Repository, cache *cache.Cache) (*service, error) {
	if repo == nil || cache == nil {
		return nil, errors.New("empty repo or cache")
	}
	return &service{
		repo:  repo,
		cache: cache,
	}, nil
}
