package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Insert(collection string, data any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("admin").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)

	return err
}

func FindAll(collection string, s interface{}) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("admin").Collection(collection)
	cursor, err := c.Find(ctx, bson.D{})

	if err != nil {
		return err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, s); err != nil {
		return err
	}

	return nil
}
