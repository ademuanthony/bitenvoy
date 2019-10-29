package handler

import (
	"context"
	"fmt"
	"github.com/ademuanthony/bitenvoy/airtime/proto/airtime"
	"github.com/gofrs/uuid"
	"time"

	orders "github.com/ademuanthony/bitenvoy/orders/proto/orders"
)

type orderHandler struct{
	store DataStore
	airtimeService airtime.AirtimeService
}

func NewOrderHandler(store DataStore) *orderHandler {
	return &orderHandler{store: store}
}

type DataStore interface {
	CreateOrder(ctx context.Context, order orders.Order) error
}

func (o orderHandler) PlaceOrder(ctx context.Context, req *orders.PlaceOrderRequest, resp *orders.PlaceOrderResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("unable to generate uuid, %s", err.Error())
	}

	order := orders.Order{
		Id:       id.String(),
		Username: req.Username,
		Product:  req.Product,
		Amount:   req.Amount,
		Date:     time.Now().UTC().Unix(),
		Details:  req.Details,
	}

	if err = o.store.CreateOrder(ctx, order); err != nil {
		return err
	}

	resp.OrderId = id.String()
	return nil
}

func (o orderHandler) MarkOrderAsPaid(context.Context, *orders.MarkOrderAsPaidRequest, *orders.EmptyMessage) error {
	panic("implement me")
}

func (o orderHandler) GetOrders(context.Context, *orders.GetOrdersRequest, *orders.GetOrdersResponse) error {
	panic("implement me")
}

func (o orderHandler) GetOrder(context.Context, *orders.GetOrderRequest, *orders.GetOrderResponse) error {
	panic("implement me")
}

