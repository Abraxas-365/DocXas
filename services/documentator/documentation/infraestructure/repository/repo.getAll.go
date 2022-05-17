package repository

import (
	"context"
	"documentator/documentation/core/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *mongoRepository) GetAll() ([]*models.Documentation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()
	collection := repo.client.Database(repo.database).Collection(repo.collection)
	documentations := []*models.Documentation{}
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &documentations); err != nil {
		return nil, err
	}
	return documentations, nil
}
