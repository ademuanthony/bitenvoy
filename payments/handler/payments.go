package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	payments "github.com/ademuanthony/bitenvoy/payments/proto/payments"
)

type Payments struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Payments) Call(ctx context.Context, req *payments.Request, rsp *payments.Response) error {
	log.Log("Received Payments.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Payments) Stream(ctx context.Context, req *payments.StreamingRequest, stream payments.Payments_StreamStream) error {
	log.Logf("Received Payments.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&payments.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Payments) PingPong(ctx context.Context, stream payments.Payments_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&payments.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
