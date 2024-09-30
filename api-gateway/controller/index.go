package controller

import (
	"go-micro.dev/v4"
	pbCards "yuno-cards/cards/proto"
	pbAuth "yuno-cards/auth/proto"

)


type Controller struct {
	ServiceMicro micro.Service
}

func (c *Controller) NewCardsController() *CardsController{
	return &CardsController{
		Cards: pbCards.NewCardsService("cards", c.ServiceMicro.Client()),
	}
}

func (c *Controller) NewAuthController() *AuthController{
	return &AuthController{
		Auth: pbAuth.NewAuthService("auth", c.ServiceMicro.Client()),
	}
}