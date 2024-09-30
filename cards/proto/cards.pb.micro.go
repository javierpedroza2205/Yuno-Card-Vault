// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cards.proto

package cards

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Cards service

func NewCardsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Cards service

type CardsService interface {
	RegisterCard(ctx context.Context, in *CreateCardRequest, opts ...client.CallOption) (*CreateCardResponse, error)
	GetCardDetails(ctx context.Context, in *GetSingleCardRequest, opts ...client.CallOption) (*InformationCard, error)
	// rpc GetCardsbyClient (GetCardsbyClientRequest) returns (InformationCardsClient) {} // get cards by client, can be n cards
	UpdateSingleCard(ctx context.Context, in *UpdateSingleCardRequest, opts ...client.CallOption) (*UpdateCardResponse, error)
	UpdateManyCards(ctx context.Context, in *UpdateManyCardsRequest, opts ...client.CallOption) (*UpdateManyCardsResponse, error)
	DeleteCard(ctx context.Context, in *GetSingleCardRequest, opts ...client.CallOption) (*CreateCardResponse, error)
}

type cardsService struct {
	c    client.Client
	name string
}

func NewCardsService(name string, c client.Client) CardsService {
	return &cardsService{
		c:    c,
		name: name,
	}
}

func (c *cardsService) RegisterCard(ctx context.Context, in *CreateCardRequest, opts ...client.CallOption) (*CreateCardResponse, error) {
	req := c.c.NewRequest(c.name, "Cards.RegisterCard", in)
	out := new(CreateCardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsService) GetCardDetails(ctx context.Context, in *GetSingleCardRequest, opts ...client.CallOption) (*InformationCard, error) {
	req := c.c.NewRequest(c.name, "Cards.GetCardDetails", in)
	out := new(InformationCard)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsService) UpdateSingleCard(ctx context.Context, in *UpdateSingleCardRequest, opts ...client.CallOption) (*UpdateCardResponse, error) {
	req := c.c.NewRequest(c.name, "Cards.UpdateSingleCard", in)
	out := new(UpdateCardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsService) UpdateManyCards(ctx context.Context, in *UpdateManyCardsRequest, opts ...client.CallOption) (*UpdateManyCardsResponse, error) {
	req := c.c.NewRequest(c.name, "Cards.UpdateManyCards", in)
	out := new(UpdateManyCardsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardsService) DeleteCard(ctx context.Context, in *GetSingleCardRequest, opts ...client.CallOption) (*CreateCardResponse, error) {
	req := c.c.NewRequest(c.name, "Cards.DeleteCard", in)
	out := new(CreateCardResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cards service

type CardsHandler interface {
	RegisterCard(context.Context, *CreateCardRequest, *CreateCardResponse) error
	GetCardDetails(context.Context, *GetSingleCardRequest, *InformationCard) error
	// rpc GetCardsbyClient (GetCardsbyClientRequest) returns (InformationCardsClient) {} // get cards by client, can be n cards
	UpdateSingleCard(context.Context, *UpdateSingleCardRequest, *UpdateCardResponse) error
	UpdateManyCards(context.Context, *UpdateManyCardsRequest, *UpdateManyCardsResponse) error
	DeleteCard(context.Context, *GetSingleCardRequest, *CreateCardResponse) error
}

func RegisterCardsHandler(s server.Server, hdlr CardsHandler, opts ...server.HandlerOption) error {
	type cards interface {
		RegisterCard(ctx context.Context, in *CreateCardRequest, out *CreateCardResponse) error
		GetCardDetails(ctx context.Context, in *GetSingleCardRequest, out *InformationCard) error
		UpdateSingleCard(ctx context.Context, in *UpdateSingleCardRequest, out *UpdateCardResponse) error
		UpdateManyCards(ctx context.Context, in *UpdateManyCardsRequest, out *UpdateManyCardsResponse) error
		DeleteCard(ctx context.Context, in *GetSingleCardRequest, out *CreateCardResponse) error
	}
	type Cards struct {
		cards
	}
	h := &cardsHandler{hdlr}
	return s.Handle(s.NewHandler(&Cards{h}, opts...))
}

type cardsHandler struct {
	CardsHandler
}

func (h *cardsHandler) RegisterCard(ctx context.Context, in *CreateCardRequest, out *CreateCardResponse) error {
	return h.CardsHandler.RegisterCard(ctx, in, out)
}

func (h *cardsHandler) GetCardDetails(ctx context.Context, in *GetSingleCardRequest, out *InformationCard) error {
	return h.CardsHandler.GetCardDetails(ctx, in, out)
}

func (h *cardsHandler) UpdateSingleCard(ctx context.Context, in *UpdateSingleCardRequest, out *UpdateCardResponse) error {
	return h.CardsHandler.UpdateSingleCard(ctx, in, out)
}

func (h *cardsHandler) UpdateManyCards(ctx context.Context, in *UpdateManyCardsRequest, out *UpdateManyCardsResponse) error {
	return h.CardsHandler.UpdateManyCards(ctx, in, out)
}

func (h *cardsHandler) DeleteCard(ctx context.Context, in *GetSingleCardRequest, out *CreateCardResponse) error {
	return h.CardsHandler.DeleteCard(ctx, in, out)
}
