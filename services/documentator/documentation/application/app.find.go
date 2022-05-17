package application

import "documentator/documentation/core/models"

func (app *app) Find(query models.FindDocumentation) ([]*models.Documentation, error) {
	return app.repo.Find(query)
}
