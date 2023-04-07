package utils

import (
	"context"
	"log"
	"time"

	"github.com/Amqp-prtcl/snowflakes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Collection *mongo.Collection

func AddDoc(c *mongo.Collection, doc interface{}) error {
	ctx, cancel := newSecContext(5)
	_, err := c.InsertOne(ctx, doc)
	cancel()
	//TODO
	return err
}

func GetDoc[T any](c *mongo.Collection, id snowflakes.ID) (*T, error) {
	ctx, cancel := newSecContext(5)
	res := c.FindOne(ctx, nil)
	cancel()
	if res.Err() != nil {
		//TODO
		return nil, res.Err()
	}
	var doc T
	err := res.Decode(&doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func GetDocs[T any](c *mongo.Collection, ids []snowflakes.ID) ([]*T, error) {
	ctx, cancel := newSecContext(10)
	defer cancel()

	cursor, err := c.Find(ctx, nil)
	if err != nil {
		//TODO
		return nil, err
	}
	defer cursor.Close(wc(newSecContext(5)))

	var docs []*T
	for cursor.Next(ctx) {
		var e T
		err = cursor.Decode(&e)
		if err != nil {
			//return docs, err
			continue
		}
		docs = append(docs, &e)
	}
	return docs, nil
}

func DoCallback(callback func(mongo.SessionContext) (interface{}, error)) error {
	sess, err := db.Client().StartSession(options.Session())
	if err != nil {
		return err
	}
	defer sess.EndSession(wc(newSecContext(5)))
	ctx, cancel := newSecContext(10)
	_, err = sess.WithTransaction(ctx, callback)
	cancel()
	return err
}

// WithTransactionExample is an example of using the Session.WithTransaction function.
func WithTransactionExample(ctx context.Context) error {
	// For a replica set, include the replica set name and a seedlist of the members in the URI string; e.g.
	// uri := "mongodb://mongodb0.example.com:27017,mongodb1.example.com:27017/?replicaSet=myRepl"
	// For a sharded cluster, connect to the mongos instances; e.g.
	// uri := "mongodb://mongos0.example.com:27017,mongos1.example.com:27017/"
	uri := ""

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}
	defer func() { _ = client.Disconnect(ctx) }()

	// Prereq: Create collections.
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)
	fooColl := client.Database("mydb1").Collection("foo", wcMajorityCollectionOpts)
	barColl := client.Database("mydb1").Collection("bar", wcMajorityCollectionOpts)

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.
		if _, err := fooColl.InsertOne(sessCtx, bson.D{{"abc", 1}}); err != nil {
			return nil, err
		}
		if _, err := barColl.InsertOne(sessCtx, bson.D{{"xyz", 999}}); err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}
	log.Printf("result: %v\n", result)
	return nil
}
