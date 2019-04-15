package nlu

import (
	context "context"
	"fmt"
	"log"

	"verbio/auth"

	"github.com/micro/go-micro/metadata"
)

type NLUSvc struct {
}

func (NLUSvc) Process(ctx context.Context, req *ProcessRequest, resp *ProcessResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return auth.ErrUnauthorized
	}
	user := md["user"]
	log.Printf("NLU request by user %s", user)
	resp.Data = fmt.Sprintf("NLU=(%s)", req.Text)
	return nil
}
