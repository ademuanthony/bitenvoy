package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/ademuanthony/bitenvoy/payments/handler"
	"github.com/ademuanthony/bitenvoy/payments/subscriber"

	payments "github.com/ademuanthony/bitenvoy/payments/proto/payments"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.payments"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payments.RegisterPaymentsHandler(service.Server(), new(handler.Payments))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.payments", service.Server(), new(subscriber.Payments))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.payments", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
