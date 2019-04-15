package main

import (
	"log"

	"time"

	"verbio/direct"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("direct"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	direct.RegisterDirectHandler(service.Server(), new(direct.DirectSvc))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
