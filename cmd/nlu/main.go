package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"

	"verbio/nlu"
)

func main() {
	service := micro.NewService(
		micro.Name("nlu"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	nlu.RegisterNLUHandler(service.Server(), new(nlu.NLUSvc))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
