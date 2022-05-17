package application

import "documentator/documentation/core/models"

func (app *app) GetAll() ([]*models.Documentation, error) {
	return app.repo.GetAll()
}
