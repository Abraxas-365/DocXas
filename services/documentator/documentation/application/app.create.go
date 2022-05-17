package application

import "documentator/documentation/core/models"

func (app *app) Create(url string) error {
	documentation := models.Documentation{}
	if err := documentation.Read(url); err != nil {
		return err
	}
	if err := app.service.CanCreate(documentation); err != nil {
		return err
	}
	return app.repo.Create(documentation)
}
