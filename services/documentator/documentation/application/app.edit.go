package application

import (
	"documentator/documentation/core/models"
)

func (app *app) Edit(edited models.Documentation) error {
	//reload information from git
	edited.Update(edited.Git)
	return app.repo.Edit(edited)
}
