package models

import (
	"yuno-cards/api-gateway/middleware"
	"go-micro.dev/v4"
)

type Router struct{
	ServiceMicro micro.Service
	ServiceAuth middleware.Authorization
}
