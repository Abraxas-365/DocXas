package repository

import (
	"context"
	"documentator/documentation/core/models"
)

func (repo *mongoRepository) Create(documentation models.Documentation) error {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()
	collection := repo.client.Database(repo.database).Collection(repo.collection)
	_, err := collection.InsertOne(ctx, documentation)
	if err != nil {
		return err
	}

	return nil

}
