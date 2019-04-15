package dialogmgr

import (
	context "context"
	"errors"
	"fmt"
	"log"
	"strings"

	"verbio/auth"
	"verbio/direct"
	"verbio/knowledge"
	"verbio/nlu"
	"verbio/taskmgr"

	client "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
)

var ErrUnknown = errors.New("Don't know what to do")

type DialogMgrSvc struct {
	Registry registry.Registry
}

func (d DialogMgrSvc) getBaseClient() client.Client {
	return client.NewClient(client.Registry(d.Registry))
}

func (d DialogMgrSvc) Message(ctx context.Context, req *MessageRequest, resp *MessageResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return auth.ErrUnauthorized
	}
	user := md["user"]
	log.Printf("Dialog manager request by user %s with text", user, req.Text)
	switch {
	case strings.Index(req.Text, "nlu:") == 0:
		svc := nlu.NewNLUService("nlu", d.getBaseClient())
		rsp, err := svc.Process(ctx, &nlu.ProcessRequest{Text: req.Text[4:]})
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Data = rsp.Data
		return nil
	case strings.Index(req.Text, "knowledge:") == 0:
		svc := knowledge.NewKnowledgeService("knowledge", d.getBaseClient())
		rsp, err := svc.Hello(ctx, &knowledge.HelloRequest{})
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Data = rsp.Text
		return nil
	case strings.Index(req.Text, "task:") == 0:
		svc := taskmgr.NewTaskMgrService("taskmgr", d.getBaseClient())
		rsp, err := svc.Hello(ctx, &taskmgr.HelloRequest{})
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Data = rsp.Text
		return nil
	case strings.Index(req.Text, "direct:") == 0:
		svc := direct.NewDirectService("direct", d.getBaseClient())
		rsp, err := svc.Hello(ctx, &direct.HelloRequest{})
		if err != nil {
			fmt.Println(err)
			return err
		}
		resp.Data = rsp.Text
		return nil

	}
	return ErrUnknown
}
