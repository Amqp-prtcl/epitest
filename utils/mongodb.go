package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	mongodbURL = ""
	epitestdb  = ""
)

var db *mongo.Database

func mongoSetup() error {
	ctx, cancel := newSecContext(10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURL).SetAppName("server-epitest").
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority())))
	cancel()
	if err != nil {
		return err
	}
	db = client.Database(epitestdb)
	return nil
}

func GetCollection(name string) (*mongo.Collection, error) {
	if db != nil {
		return db.Collection(name), nil
	}
	return nil, fmt.Errorf("get collection: server not connected to mongodb instance")
}

func mongoClose() error {
	ctx, cancel := newSecContext(20)
	defer cancel()
	return db.Client().Disconnect(ctx)
}
