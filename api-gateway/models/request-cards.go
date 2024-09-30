package models

import cards "yuno-cards/cards/proto"

type CreateCardRequest struct{
	CardHolderName string `json:"card_holder_name" validate:"required"`
    CardType       string `json:"card_type" validate:"required"`
    ExpiryDate     string `json:"expiry_date" validate:"required"`
    CodeCard       string `json:"code_card" validate:"required"`
    Alias          string `json:"alias" validate:"required"`
    NumberCard	   string `json:"number_card" validate:"required"`
}

type UpdateSinglecardRequest struct {
    CardHolderName string `json:"card_holder_name"`
    Alias          string `json:"alias"`
    CardID         string `json:"id_card" validate:"required"`
}

type UpdateManycardRequest struct {
    Cards []*cards.UpdateSingleCardRequest `json:"cards"`
}