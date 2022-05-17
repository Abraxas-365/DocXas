package ports

import "documentator/documentation/core/models"

type Repository interface {
	Create(models.Documentation) error
	Edit(models.Documentation) error
	Find(models.FindDocumentation) ([]*models.Documentation, error)
	GetAll() ([]*models.Documentation, error)
	DoesExist(name string, version int) bool
}
