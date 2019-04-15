package main

import (
	"log"
	"time"
	"verbio/dialogmgr"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("dialogmgr"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()
	opts := service.Options()
	dialogmgr.RegisterDialogMgrHandler(service.Server(), &dialogmgr.DialogMgrSvc{opts.Registry})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
