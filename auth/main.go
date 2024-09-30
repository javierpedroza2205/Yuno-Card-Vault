package main

import (
	"context"
	"yuno-cards/auth/handler"
	pbAuth "yuno-cards/auth/proto"
	"yuno-cards/auth/repository"

	log "github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"github.com/joho/godotenv"
	"os"

)

func main() {
	srv := micro.NewService(
		micro.Name("auth"),
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


	Auth := handler.Auth{
		ClientMongoDB: sessionMongoDB,
	}
	pbAuth.RegisterAuthHandler(srv.Server(), &Auth)

	if err := srv.Run(); err != nil {
		log.Fatal("Error while trying to run Auth service: ", err)
	}
}
