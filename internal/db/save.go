package db

import "context"

func Insert(collection string, data any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("subscribe").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)

	return err
}
