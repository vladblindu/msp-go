package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"msp-go/config"
	"time"
)

func (d *Db) createClient(conf config.Config) {
	uri := createUri(conf)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	d.Client = *client
	d.Ctx, d.cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(d.Ctx)
	if err != nil {
		d.cancel()
		log.Fatal(err)
	}
}

func (d *Db) Close() {
	err := d.Client.Disconnect(d.Ctx)
	if err != nil {
		d.cancel()
		log.Fatal(err)
	}
	d.cancel()
}

func New(conf config.Config) *Db {
	d := new(Db)
	d.createClient(conf)
	coll := d.Client.Database(conf.MspCacheCredentialsDatabase).Collection(conf.MspCacheCredentialsCollection)
	d.ContainersCollection = *coll
	return d
}
