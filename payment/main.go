package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/ademuanthony/bitenvoy/payment/handler"
	"github.com/ademuanthony/bitenvoy/payment/subscriber"

	payment "github.com/ademuanthony/bitenvoy/payment/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.payment", service.Server(), new(subscriber.Payment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
