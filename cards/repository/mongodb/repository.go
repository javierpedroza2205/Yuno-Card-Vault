package cards

import (
	"context"
	"os"
	"reflect"
	"runtime"
	"time"
	pbCards "yuno-cards/cards/proto"
	"yuno-cards/cards/repository/models"
	"yuno-cards/cards/utils"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"path/filepath"
)


const (
	dbname = "yuno"
	cardsCollection = "cards"
)

type Repository interface {
	RegisterCard(context.Context, *pbCards.CreateCardRequest)(string, error)
	GetCardDetails(context.Context, string, string)(*pbCards.InformationCard, error)
	DeleteCard(context.Context, string, string)(error, int)
	UpdateSingleCard(context.Context, *pbCards.UpdateSingleCardRequest)(*pbCards.UpdateCardResponse, error)
	UpdateManyCards(context.Context, *pbCards.UpdateManyCardsRequest)([]*pbCards.UpdateCardResponse, error)

}

type CardsRepository struct {
	Client *mongo.Client
}

func (repo *CardsRepository) collectionCards() *mongo.Collection{
	return repo.Client.Database(dbname).Collection(cardsCollection)
}


func (repo *CardsRepository) RegisterCard(ctx context.Context, req *pbCards.CreateCardRequest) (string, error){
	now := time.Now().String()
	idCard := primitive.NewObjectID()
	var newCard = &models.CardData{
		ID: idCard,
		CreatedAt: now,
		CardType: req.CardType,
		CardHoldername: req.CardHolderName,
		ExpiryDate: req.ExpiryDate,
		CodeCard: req.CodeCard,
		Alias: req.Alias,
		CardNumber: req.NumberCard,
		IDClient: req.IdClient,
	}

	keyValue := GetKey()

	v := reflect.ValueOf(newCard).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		if field.Type().Kind() == reflect.String && fieldType.Name != "IDClient" {
			newValue, _ := utils.Encrypt(field.String(), keyValue )
			field.SetString(newValue)
		}
	}
	

	_, errCreate := repo.collectionCards().InsertOne(context.Background(), newCard)
	if errCreate != nil {
		return "", errCreate
	}
	return  idCard.Hex(), nil
}


func (repo *CardsRepository) GetCardDetails(ctx context.Context, IdCard string, IDClient string) (*pbCards.InformationCard, error){
	var result models.CardDataResponse
	objectID, errCast := primitive.ObjectIDFromHex(IdCard)
	if errCast != nil {
		return nil, errCast
	}
	filter := bson.D{{"_id", objectID}, {"id_client", IDClient}}

	err := repo.collectionCards().FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}

	keyValue := GetKey()


	v := reflect.ValueOf(&result).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Type().Kind() == reflect.String {
			newValue, _ := utils.Decrypt(field.String(), keyValue )
			field.SetString(newValue)
		}
	}

	// map result to proto
	res := &pbCards.InformationCard{
		CardHolderName: result.CardHoldername,
		CardType: result.CardType,
		Alias: result.Alias,
		NumberCard: result.NumberCard[len(result.NumberCard)-4:],
		CreatedAt: result.CreatedAt,
	}
	return res, nil

}

func (repo *CardsRepository) DeleteCard(ctx context.Context, IdCard string, IDClient string) (error, int){
	objectID, errCast := primitive.ObjectIDFromHex(IdCard)
	if errCast != nil {
		return errCast, 0
	}
	filter := bson.D{{"_id", objectID}, {"id_client", IDClient}}

	res , err := repo.collectionCards().DeleteOne(ctx, filter)
	if err != nil {
		return err, 0
	}

	if res.DeletedCount != 1 {
		return nil, 0
	}else{
		return nil, 1
	}
}

func (repo *CardsRepository) UpdateSingleCard(ctx context.Context, req *pbCards.UpdateSingleCardRequest) (*pbCards.UpdateCardResponse, error){
	response := &pbCards.UpdateCardResponse{}
	updateAux := bson.M{}
	var updateCard = &models.CardDataUpdate{
		Card_holder_name: req.CardHolderName,
		Alias: req.Alias,
	}
	keyValue := GetKey()
	v := reflect.ValueOf(updateCard).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if field.Type().Kind() == reflect.String && field.String() != "" {
			newValue, _ := utils.Encrypt(field.String(), keyValue )
			updateAux[fieldType.Name] = newValue
		}
		
	}
	objectID, errCast := primitive.ObjectIDFromHex(req.IdCard)
	if errCast != nil {
		response = &pbCards.UpdateCardResponse{
			Status: "406",
			Message: "Error Update Card, Cause Cast ID",
			IdCard: req.IdCard,
		}
		return response, errCast
	}
	filter := bson.D{{"_id", objectID}, {"id_client", req.IdClient}}
	update := bson.M{
		"$set": updateAux,
	}

	resUpdate, errupdate := repo.collectionCards().UpdateOne(ctx, filter, update)
	if errupdate != nil {
		response = &pbCards.UpdateCardResponse{
			Status: "406",
			Message: "Error Update Card, Cause Error Query",
			IdCard: req.IdCard,
		}
		return response, errCast
	}

	if resUpdate.ModifiedCount != 1{
		response = &pbCards.UpdateCardResponse{
			Status: "406",
			Message: "Card Doesnt Exist",
			IdCard: req.IdCard,
		}

	}else{
		response = &pbCards.UpdateCardResponse{
			Status: "200",
			Message: "Card Update Success",
			IdCard: req.IdCard,
		}
	}


	return response, nil
}

func (repo *CardsRepository) UpdateManyCards(ctx context.Context, req *pbCards.UpdateManyCardsRequest) ([]*pbCards.UpdateCardResponse, error){
	var responseReturn []*pbCards.UpdateCardResponse
	internalRequest := &pbCards.UpdateSingleCardRequest{}
	for _ , value := range req.InformationUpdate{
		internalRequest = &pbCards.UpdateSingleCardRequest{
			CardHolderName: value.CardHolderName,
			Alias: value.Alias,
			IdCard: value.IdCard,
			IdClient: req.IdClient,
		}

		responseInternal, _:= repo.UpdateSingleCard(ctx, internalRequest)
		responseReturn = append(responseReturn, responseInternal)
	}
	return responseReturn, nil
}

func GetKey()(string){

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	err := godotenv.Load(basepath + "/../../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv("KEY")	
	return key
}