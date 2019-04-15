package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"

	"verbio/auth"
)

func main() {
	service := micro.NewService(
		micro.Name("auth"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	auth.RegisterAuthHandler(service.Server(), new(auth.AuthSvc))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
