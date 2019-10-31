package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	payments "github.com/ademuanthony/bitenvoy/payments/proto/payments"
)

type Payments struct{}

func (e *Payments) Handle(ctx context.Context, msg *payments.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *payments.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
