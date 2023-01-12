package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	test66 "test66/proto/test66"
)

type Test66 struct{}

func (e *Test66) Handle(ctx context.Context, msg *test66.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test66.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
