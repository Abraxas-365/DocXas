package application

import (
	"documentator/documentation/core/models"
	"documentator/documentation/core/ports"
	"documentator/documentation/core/service"
)

type Application interface {
	Create(url string) error
	Edit(models.Documentation) error
	Find(models.FindDocumentation) ([]*models.Documentation, error)
	GetAll() ([]*models.Documentation, error)
}

type app struct {
	repo    ports.Repository
	service service.Service
}

func NewApp(repo ports.Repository, service service.Service) Application {
	return &app{
		repo,
		service,
	}
}
