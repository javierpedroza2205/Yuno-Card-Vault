package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type CardData struct{
	ID             primitive.ObjectID ` bson:"_id"`
	CreatedAt      string        ` bson:"created_at"`
	CardType		   string             ` bson:"card_type"`
	CardHoldername	   string			  ` bson:"card_holder_name"`
	ExpiryDate	       string			  ` bson:"expiry_date"`
	CodeCard		   string			  ` bson:"code_card"`
	Alias		       string			  ` bson:"alias"`
	CardNumber         string      `bson:"number_card"`
	IDClient           string      `bson:"id_client"`

}

// return card strcture

type CardDataResponse struct {
	ID  string `bson:"_id"`
	CardType		   string             `bson:"card_type"`
	CardHoldername	   string			  `bson:"card_holder_name"`
	Alias		       string			  `bson:"alias"`
	NumberCard         string      `bson:"number_card"`
	CreatedAt          string        `bson:"created_at"`
}

// update 

type CardDataUpdate struct {
	Card_holder_name	   string			  `bson:"card_holder_name"`
	Alias		       string			  `bson:"alias"`
}