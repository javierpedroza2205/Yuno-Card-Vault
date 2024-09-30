package auth

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"time"
	pbAuth "yuno-cards/auth/proto"

	"yuno-cards/auth/repository/models"
	"yuno-cards/auth/utils"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


const (
	dbname = "yuno"
	authCollection = "auth"
	
)


type Repository interface {
	DoAuth(context.Context, *pbAuth.AuthRequest)(*pbAuth.AuthResponse, error)
	RegisterUser(context.Context, *pbAuth.AuthRequest)(error)


}

type AuthRepository struct {
	Client *mongo.Client
}


func (repo *AuthRepository) collectionAuth() *mongo.Collection{
	return repo.Client.Database(dbname).Collection(authCollection)
}

func (a *AuthRepository) DoAuth(ctx context.Context, req *pbAuth.AuthRequest ) (*pbAuth.AuthResponse, error){
	var result models.GenericUser
	response := &pbAuth.AuthResponse{}

	filter := bson.D{{"username", req.Username}}
	err := a.collectionAuth().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		response = &pbAuth.AuthResponse{
			Status: "403",
			Message: "Error Find UserName",
		}
		return response, err
	}
	isMatch := utils.ComparePasswords(result.Password, req.Password)
    if isMatch {
		tokenValue, errtoken := a.GenerateToken(ctx, result.ID.Hex())
		if errtoken != nil{
			response = &pbAuth.AuthResponse{
				Status: "403",
				Message: "Error Generate Token",
			}
			return response, errtoken
		}
		update := bson.M{
			"$set": bson.M{
				"last_login": time.Now(),
			},
		}
		_, errUpdate := a.collectionAuth().UpdateOne(context.Background(), filter, update)
		if errUpdate != nil{
			response = &pbAuth.AuthResponse{
				Status: "403",
				Message: "Error Generate Token",
			}
			return response, errUpdate
		}
		response = &pbAuth.AuthResponse{
			Status: "200",
			Message: "Auth Success",
			Token: tokenValue,
		}
		return response, nil


    } else {
		response = &pbAuth.AuthResponse{
			Status: "403",
			Message: "Error in Password",
		}
		return response, err
    }
}

func (a *AuthRepository) RegisterUser(ctx context.Context, req *pbAuth.AuthRequest) (error){
	hashedPassword, errHashes := utils.HashPassword(req.Password)
    if errHashes != nil {
        return errHashes
    }

	var newUser = &models.GenericUser{
		ID: primitive.NewObjectID(),
		Username: req.Username,
		Password: hashedPassword,
	} 

	_, errCreate := a.collectionAuth().InsertOne(context.Background(), newUser)
	if errCreate != nil {
		return errCreate
	}
	
	return  nil
	
}

func (a *AuthRepository) GenerateToken(ctx context.Context, userID string) (string, error){
	keyAuth, kid := GetKeys()
	token, err := utils.GenerateJWT(userID, []byte(keyAuth), kid)
	if err != nil{
		return "", err
	}
	return token, nil
}


func GetKeys()(string, string){

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	keyAuth := os.Getenv("KEY_AUTH")	
	kID := os.Getenv("KID")
	return keyAuth, kID
}