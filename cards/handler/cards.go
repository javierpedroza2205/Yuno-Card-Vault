package handler

import (
	"context"
	pbCards "yuno-cards/cards/proto"
	db "yuno-cards/cards/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"

)


type YunoCards struct {
	ClientMongoDB *mongo.Client
}

func (y *YunoCards) GetClient() db.Repository {
	return &db.CardsRepository{
		Client: y.ClientMongoDB,
	}
}

func (y *YunoCards) RegisterCard(ctx context.Context, req *pbCards.CreateCardRequest,  res *pbCards.CreateCardResponse) error{
	err := y.GetClient().RegisterCard(ctx, req)
	if err != nil{
		res.Message = "Error Create Card"
		res.Status = "406"
	}else{
		res.Message = "Create Card Success"
		res.Status = "200"
	}
	return nil
}

func (y *YunoCards) GetCardDetails(ctx context.Context, req *pbCards.GetSingleCardRequest, res *pbCards.InformationCard) error {
	infoCard, err := y.GetClient().GetCardDetails(ctx, req.IdCard, req.IdClient)
	if err != nil {
		return err
	}else{
		res.Alias = infoCard.Alias
		res.CardHolderName = infoCard.CardHolderName
		res.NumberCard = infoCard.NumberCard
		res.CardType = infoCard.CardType
		res.IdCard = infoCard.IdCard
		res.CreatedAt = infoCard.CreatedAt
	}
	return err
	
}

func (y *YunoCards) UpdateSingleCard(ctx context.Context, req *pbCards.UpdateSingleCardRequest, res *pbCards.UpdateCardResponse) error {
	response, _ := y.GetClient().UpdateSingleCard(ctx, req)
	res.IdCard = response.IdCard
	res.Message = response.Message
	res.Status = response.Status
	return nil
}

func (y *YunoCards) UpdateManyCards(ctx context.Context, req *pbCards.UpdateManyCardsRequest, res *pbCards.UpdateManyCardsResponse) error {
	response, _ := y.GetClient().UpdateManyCards(ctx, req)
	res.Process = response
	return nil
}

func (y *YunoCards) DeleteCard(ctx context.Context, req *pbCards.GetSingleCardRequest, res *pbCards.CreateCardResponse) error {
	err, delElements := y.GetClient().DeleteCard(ctx, req.IdCard, req.IdClient)
	if err != nil{
		res.Message = "Error Delete Card"
		res.Status = "406"
	}else{
		if delElements == 1{
			res.Message = "Delete Card Success"
			res.Status = "200"
		}else{
			res.Message = "Card Doesnt Exist"
			res.Status = "406"
		}
	}

	return nil
}