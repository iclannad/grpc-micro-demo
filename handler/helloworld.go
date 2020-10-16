package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/util/log"

	helloworld "grpc-micro-demo/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	log.Log("Received Helloworld.Call request")

	time.Sleep(30 * time.Second)

	rsp.Msg = "Hello " + req.Name
	return nil
}
