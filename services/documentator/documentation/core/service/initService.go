package service

import (
	"documentator/documentation/core/models"
	"documentator/documentation/core/ports"
	"errors"
)

var (
	ErrorExist = errors.New("Service already exists")
)

type Service interface {
	CanCreate(models.Documentation) error
}

type service struct {
	repo ports.Repository
}

func NewService(repo ports.Repository) Service {
	return &service{
		repo,
	}
}
