package application

import (
	"documentator/documentation/core/models"
)

func (app *app) Edit(edited models.Documentation) error {
	return app.repo.Edit(edited)
}
