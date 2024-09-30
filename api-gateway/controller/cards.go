package controller

import (
	"context"
	"fmt"
	"net/http"
	"yuno-cards/api-gateway/models"
	"yuno-cards/api-gateway/middleware"
	microErrors "yuno-cards/api-gateway/utils"
	pbCards "yuno-cards/cards/proto"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CardsController struct {
	Cards pbCards.CardsService
}

const (
	createCardEvents = "yuno.services.create.card"
	updateCardEvents = "yuno.services.update.single.card"
	getSinclgeCardEvents = "yuno.services.get.info.card"

)

func (cards *CardsController) HandlerRegisterCard(c echo.Context) error {
	bodyPayload := new(models.CreateCardRequest)
	
	if err := c.Bind(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        createCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[1].",
				Code:        "create_card_bad_request",
			},
		}
		log.Error("Error ocurred while Bind CreateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := c.Validate(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        createCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[2].",
				Code:        "create_card_bad_request",
			},
		}
		log.Error("Error ocurred while validate CreateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	id_client := middleware.GetUserToken(c, "_id")

	response, errApi := cards.Cards.RegisterCard(context.Background(), &pbCards.CreateCardRequest{
		CardHolderName: bodyPayload.CardHolderName,
		CardType: bodyPayload.CardType,
		ExpiryDate: bodyPayload.ExpiryDate,
		CodeCard: bodyPayload.CodeCard,
		NumberCard: bodyPayload.NumberCard,
		Alias: bodyPayload.Alias,
		IdClient: id_client,
	})

	if errApi != nil{
		fmt.Println("error create card api", errApi.Error())
		return c.JSON(http.StatusInternalServerError, errApi.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (cards *CardsController) HandlerGetSingleCard(c echo.Context) error {
	cardID := c.Param("cardId")
	id_client := middleware.GetUserToken(c, "_id")
	rsp, err := cards.Cards.GetCardDetails(context.Background(), &pbCards.GetSingleCardRequest{
		IdCard: cardID,
		IdClient: id_client,
	})

	if err != nil{
		fmt.Println("error get card api", err.Error())
		if err.Error() == "mongo: no documents in result"{
			errorResponse := microErrors.Error{
				Status: http.StatusText(200),
				Error: microErrors.ErrorReason{
					Type:        getSinclgeCardEvents,
					StatusCode:  http.StatusOK,
					Message:     err.Error(),
					UserMessage: "Error find card by user.",
					Code:        "get_info_card_bad_request",
				},
			}
			return c.JSON(http.StatusOK, errorResponse)
		}
		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusOK, rsp)
}

func (cards *CardsController) HandlerDeleteCard(c echo.Context) error {
	cardID := c.Param("cardId")
	id_client := middleware.GetUserToken(c, "_id")
	rsp, err := cards.Cards.DeleteCard(context.Background(), &pbCards.GetSingleCardRequest{
		IdCard: cardID,
		IdClient: id_client,
	})

	if err != nil{
		fmt.Println("error get card api", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rsp)
}

func (cards *CardsController) HandlerUpdateSingleCard(c echo.Context) error {
	bodyPayload := new(models.UpdateSinglecardRequest)
	if err := c.Bind(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        updateCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[1].",
				Code:        "update_card_bad_request",
			},
		}
		log.Error("Error ocurred while Bind UpdateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := c.Validate(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        updateCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[2].",
				Code:        "update_card_bad_request",
			},
		}
		log.Error("Error ocurred while validate UpdateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	id_client := middleware.GetUserToken(c, "_id")


	response, errApi := cards.Cards.UpdateSingleCard(context.Background(), &pbCards.UpdateSingleCardRequest{
		CardHolderName: bodyPayload.CardHolderName,
		Alias: bodyPayload.Alias,
		IdCard: bodyPayload.CardID, // required
		IdClient: id_client,
	})

	if errApi != nil {
		fmt.Println("error update card api", errApi.Error())
		return c.JSON(http.StatusInternalServerError, errApi.Error())

	}
	return c.JSON(http.StatusOK, response)
}

func (cards *CardsController) HandlerUpdateManyCards(c echo.Context) error {
	bodyPayload := new(models.UpdateManycardRequest)
	if err := c.Bind(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        updateCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[1].",
				Code:        "update_card_bad_request",
			},
		}
		log.Error("Error ocurred while Bind UpdateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := c.Validate(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        updateCardEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[2].",
				Code:        "update_card_bad_request",
			},
		}
		log.Error("Error ocurred while validate UpdateCard request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	id_client := middleware.GetUserToken(c, "_id")
	response, errApi := cards.Cards.UpdateManyCards(context.Background(), &pbCards.UpdateManyCardsRequest{
		InformationUpdate: bodyPayload.Cards ,
		IdClient: id_client,
		
	})

	if errApi != nil {
		fmt.Println("error update manycards api", errApi.Error())
		return c.JSON(http.StatusInternalServerError, errApi.Error())


	}
	return c.JSON(http.StatusOK, response)
}