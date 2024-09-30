package main

import (
	"context"
	"os"
	"yuno-cards/cards/handler"
	pbCards "yuno-cards/cards/proto"
	"yuno-cards/cards/repository"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go-micro.dev/v4"
)

func main() {
	srv := micro.NewService(
		micro.Name("cards"),
		micro.Version("v1"),
	)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	sessionMongoDB, err := repository.ConnectMongo(dbHost)
	defer sessionMongoDB.Disconnect(context.Background())
	if err != nil {
		log.Panic("Could not connect to datastore ", err)
	}

	Cards := handler.YunoCards{
		ClientMongoDB: sessionMongoDB,
	}
	pbCards.RegisterCardsHandler(srv.Server(), &Cards)

	if err := srv.Run(); err != nil {
		log.Fatal("Error while trying to run Cards service: ", err)
	}
}



