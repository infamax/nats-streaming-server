package service

import (
	"errors"
	"github.com/infamax/nats-streaming-server/internal/cache"
	"github.com/infamax/nats-streaming-server/internal/models"
	"github.com/infamax/nats-streaming-server/internal/repository"
)

type Service interface {
	AddModelDB(order *models.Order) error
	GetModelDB(uuid string) (*models.Order, error)
	UpdateModelDB(order *models.Order) error
	DeleteModelDB(uuid string) error
	GetAllModels() ([]models.Order, error)
	AddModelCache(order *models.Order) error
	GetModelCache(uuid string) (*models.Order, error)
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
