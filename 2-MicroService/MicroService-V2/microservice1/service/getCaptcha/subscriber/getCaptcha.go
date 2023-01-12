package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	getCaptcha "bj38web/service/getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Handle(ctx context.Context, msg *getCaptcha.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *getCaptcha.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
