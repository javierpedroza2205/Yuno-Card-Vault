package main

import (
	router "yuno-cards/api-gateway/server"

	"go-micro.dev/v4/web"
)
func main()  {

	opts := []web.Option{
		web.Name("yuno.service.api"),
		web.Address(":8080"),
	}

	service := web.NewService(opts...)

	router := router.Router{
		ServiceMicro: service.Options().Service,
	}

	e := router.NewRouter()
	service.Handle("/", e)	
	service.Run()
	
}