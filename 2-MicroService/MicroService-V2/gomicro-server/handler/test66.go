package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	test66 "test66/proto/test66"
)

type Test66 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test66) Call(ctx context.Context, req *test66.Request, rsp *test66.Response) error {
	log.Log("Received Test66.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test66) Stream(ctx context.Context, req *test66.StreamingRequest, stream test66.Test66_StreamStream) error {
	log.Logf("Received Test66.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test66.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test66) PingPong(ctx context.Context, stream test66.Test66_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test66.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
