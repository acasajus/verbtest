package nlu

import (
	"context"
	"fmt"
	"log"
	"testing"

	micro "github.com/micro/go-micro"
	client "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/memory"
)

var gSvc micro.Service
var gReg registry.Registry

func init() {
	gReg = memory.NewRegistry()
	gSvc = micro.NewService(
		micro.Name("nlu-test"),
		micro.Version("latest"),
		micro.Registry(gReg),
	)

	gSvc.Init()

	RegisterNLUHandler(gSvc.Server(), new(NLUSvc))

	gSvc.Init()
	go func() {
		if err := gSvc.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	//time.Sleep(time.Second)
	w, err := gReg.Watch()
	if err != nil {
		panic(err)
	}
	_, err = w.Next()
	if err != nil {
		panic(err)
	}
}

func getBaseClient() client.Client {
	return client.NewClient(client.Registry(gReg))
}

func TestNLUServer(t *testing.T) {
	client := NewNLUService("nlu-test", getBaseClient())
	req := &ProcessRequest{Text: "hello"}
	resp, err := client.Process(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	expected := fmt.Sprintf("NLU=(%s)", req.Text)
	if resp.Data != expected {
		t.Errorf("Unexpected response: %s vs %s", expected, resp.Data)
	}
}
