package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	mygomicroserver "my-gomicro-server/proto/my-gomicro-server"
)

type MyGomicroServer struct{}

func (e *MyGomicroServer) Handle(ctx context.Context, msg *mygomicroserver.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *mygomicroserver.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
