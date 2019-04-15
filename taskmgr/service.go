package taskmgr

import (
	context "context"
	"fmt"
	"log"

	"verbio/auth"

	"github.com/micro/go-micro/metadata"
)

type TaskMgrSvc struct {
}

func (TaskMgrSvc) Hello(ctx context.Context, req *HelloRequest, resp *HelloResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return auth.ErrUnauthorized
	}
	user := md["user"]
	resp.Text = fmt.Sprintf("TaskMgr request by user %s", user)
	log.Println(resp.Text)
	return nil
}
