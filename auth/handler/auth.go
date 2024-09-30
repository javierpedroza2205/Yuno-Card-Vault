package handler

import (
	"context"
	pbAuth "yuno-cards/auth/proto"
	db "yuno-cards/auth/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	ClientMongoDB *mongo.Client
}

func (a *Auth) GetClient() db.Repository {
	return &db.AuthRepository{
		Client: a.ClientMongoDB,
	}
}


func (a *Auth) DoAuth(ctx context.Context, req *pbAuth.AuthRequest, res *pbAuth.AuthResponse) error{
	authResponse, err := a.GetClient().DoAuth(ctx,req)
	if err != nil{
		res.Message = authResponse.Message
		res.Status = authResponse.Status
	}else{
		res.Message = authResponse.Message
		res.Status = authResponse.Status
		res.Token = authResponse.Token
	}
	return nil
}

func (a *Auth) RegisterUser(ctx context.Context, req *pbAuth.AuthRequest, res *pbAuth.AuthResponse) error{
	err := a.GetClient().RegisterUser(ctx,req)
	if err != nil{
		res.Message = "Error Register User"
		res.Status = "406"
	}else{
		res.Message = "Register User Success"
		res.Status = "200"
	}
	

	return nil
}