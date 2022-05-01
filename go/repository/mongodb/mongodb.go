package repository

import (
	"context"
	"log"
	"time"

	"github.com/renatoviolin/go-node/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, timeout int) (*mongo.Client, error) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewMongoRepository(mongoURL, mongoDB string, timeout int) (*MongoRepository, error) {
	repo := &MongoRepository{
		database: mongoDB,
		timeout:  time.Duration(timeout) * time.Second,
	}

	client, err := newMongoClient(mongoURL, timeout)
	if err != nil {
		return nil, err
	}

	repo.client = client
	return repo, nil
}

func (m *MongoRepository) FindAll() (*[]dto.Payload, error) {
	ctx := context.Background()
	collection := m.client.Database(m.database).Collection("sample")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("error: %+v\n", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var results []dto.Payload
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return &results, nil
}
