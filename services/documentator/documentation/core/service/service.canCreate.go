package service

import "documentator/documentation/core/models"

func (service *service) CanCreate(documentation models.Documentation) error {
	if service.repo.DoesExist(documentation.ServiceName, documentation.Version) {
		return ErrorExist
	}
	return nil
}
