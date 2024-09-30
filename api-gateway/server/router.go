package router

import (
	"net/http"
	"yuno-cards/api-gateway/controller"
	"yuno-cards/api-gateway/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)


type Router models.Router
func (s *Router) NewRouter() *echo.Echo{
	e := echo.New()
	APIGENERIC := e.Group("/api/")
	APIACARDS := e.Group("/api/v1/yuno/cards/")
	APIAUTH := e.Group("/api/v1/yuno/auth/")

	initController := controller.Controller{
		ServiceMicro: s.ServiceMicro,
	}

	e.Validator = &models.CustomValidator{Validator: validator.New()}

	var (
		scvCardsController   = initController.NewCardsController()
		scvAuthController    = initController.NewAuthController()
	)
	
	APIGENERIC.GET("internal/healt", healt)

	// Cards endpoints
	APIACARDS.POST("create", scvCardsController.HandlerRegisterCard, s.ServiceAuth.Authorize())
	APIACARDS.GET("getCard/:cardId", scvCardsController.HandlerGetSingleCard, s.ServiceAuth.Authorize())
	//API.GET("getCardsbyClient", scvCardsController.HandlerGetCardsbyClient)
	APIACARDS.DELETE("delete/:cardId", scvCardsController.HandlerDeleteCard, s.ServiceAuth.Authorize())
	APIACARDS.PUT("updateSingleCard", scvCardsController.HandlerUpdateSingleCard, s.ServiceAuth.Authorize())
	APIACARDS.POST("updateManyCards", scvCardsController.HandlerUpdateManyCards, s.ServiceAuth.Authorize())
	// extra endpoint card
	//API.POST("rotate-keys", )
	// Auth endpoints
	APIAUTH.POST("login", scvAuthController.HandlerAuth)
	APIAUTH.POST("register", scvAuthController.HandlerRegister)

	return e

}

func healt(c echo.Context) error {
	return c.String(http.StatusOK, "")
}