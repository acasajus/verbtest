package main

import (
	"log"
	"time"

	"verbio/knowledge"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("knowledge"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	knowledge.RegisterKnowledgeHandler(service.Server(), new(knowledge.KnowledgeSvc))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
