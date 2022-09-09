package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func (d *Db) GetContainers() []Container {
	cursor, err := d.ContainersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Print("No containers found")
			return []Container{}
		}
		panic(err)
	}
	var results []Container
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}
