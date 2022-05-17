package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *mongoRepository) DoesExist(name string, version int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()
	collection := repo.client.Database(repo.database).Collection(repo.collection)

	filter := bson.M{"name": name, "version": version}

	check := collection.FindOne(ctx, filter)

	if check != nil {
		return true
	}
	return false

}
