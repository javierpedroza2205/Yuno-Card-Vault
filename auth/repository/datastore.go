package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(host string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(host)
	clientOptions.SetMaxPoolSize(500)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB Services!")

	return client, nil
}