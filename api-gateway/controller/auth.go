package controller

import (
	"context"
	"fmt"
	"net/http"
	"yuno-cards/api-gateway/models"
	microErrors "yuno-cards/api-gateway/utils"
	pbAuth "yuno-cards/auth/proto"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

const (
	authEvents = "yuno.services.auth"
)

type AuthController struct {
	Auth pbAuth.AuthService
}

func (auth *AuthController) HandlerAuth(c echo.Context) error {
	bodyPayload := new(models.AuthRequest)
	

	if err := c.Bind(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        authEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[1].",
				Code:        "auth_user_bad_request",
			},
		}
		log.Error("Error ocurred while Bind AuthUser request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := c.Validate(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        authEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[2].",
				Code:        "auth_user_bad_request",
			},
		}
		log.Error("Error ocurred while validate AuthUser request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	response, err := auth.Auth.DoAuth(context.Background(), &pbAuth.AuthRequest{
		Username: bodyPayload.Username,
		Password: bodyPayload.Password,
	})
	if err != nil{
		fmt.Println("error auth api", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}

func (auth *AuthController) HandlerRegister(c echo.Context) error{
	bodyPayload := new(models.AuthRequest)
	if err := c.Bind(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        authEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[1].",
				Code:        "register_user_bad_request",
			},
		}
		log.Error("Error ocurred while Bind RegisterUser request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if err := c.Validate(bodyPayload); err != nil {
		errorResponse := microErrors.Error{
			Status: http.StatusText(400),
			Error: microErrors.ErrorReason{
				Type:        authEvents,
				StatusCode:  http.StatusBadRequest,
				Message:     err.Error(),
				UserMessage: "Bad request: Please inspect the information and try again[2].",
				Code:        "register_user_bad_request",
			},
		}
		log.Error("Error ocurred while validate RegisterUser request: ", errorResponse)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	response, err := auth.Auth.RegisterUser(context.Background(), &pbAuth.AuthRequest{
		Username: bodyPayload.Username,
		Password: bodyPayload.Password,
	})
	if err != nil{
		fmt.Println("error authregister api", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}