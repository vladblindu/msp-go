package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Db struct {
	Client               mongo.Client
	Ctx                  context.Context
	cancel               func()
	ContainersCollection mongo.Collection
}

type Container struct {
	Id       string `json:"_id"`
	Kind     string `json:"kind"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Docker   string `json:"docker"`
}
