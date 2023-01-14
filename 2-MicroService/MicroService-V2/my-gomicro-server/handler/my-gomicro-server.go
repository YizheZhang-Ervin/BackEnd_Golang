package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	mygomicroserver "my-gomicro-server/proto/my-gomicro-server"
)

type MyGomicroServer struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MyGomicroServer) Call(ctx context.Context, req *mygomicroserver.Request, rsp *mygomicroserver.Response) error {
	log.Info("Received MyGomicroServer.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MyGomicroServer) Stream(ctx context.Context, req *mygomicroserver.StreamingRequest, stream mygomicroserver.MyGomicroServer_StreamStream) error {
	log.Infof("Received MyGomicroServer.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&mygomicroserver.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MyGomicroServer) PingPong(ctx context.Context, stream mygomicroserver.MyGomicroServer_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&mygomicroserver.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
