package repository

import (
	"context"
	"documentator/documentation/core/models"
	"errors"
)

func (repo *mongoRepository) Edit(documentation models.Documentation) error {
	filterQuery := make(map[string]interface{})
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()
	collection := repo.client.Database(repo.database).Collection(repo.collection)
	switch true {
	case len(documentation.Id) != 0:
		filterQuery["_id"] = documentation.Id
	case len(documentation.ServiceName) != 0 && documentation.Version != 0:
		filterQuery["service_name"] = documentation.ServiceName
		filterQuery["version"] = documentation.Version
	}

	updated := collection.FindOneAndUpdate(ctx, filterQuery, documentation)
	if updated == nil {
		return errors.New("document doesnt exist")
	}
	return nil

}
