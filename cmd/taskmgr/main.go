package main

import (
	"log"

	"time"

	"verbio/taskmgr"

	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("taskmgr"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()

	taskmgr.RegisterTaskMgrHandler(service.Server(), new(taskmgr.TaskMgrSvc))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
