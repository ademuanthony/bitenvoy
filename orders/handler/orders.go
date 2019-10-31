package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"strings"
	"time"

	"github.com/ademuanthony/bitenvoy/airtime/proto/airtime"
	orders "github.com/ademuanthony/bitenvoy/orders/proto/orders"
	"github.com/gofrs/uuid"
)

const (
	Airtime = "airtime"
	Data = "data"
)

type orderHandler struct{
	store DataStore
	airtimeService airtime.AirtimeService
}

func NewOrderHandler(store DataStore, airtimeService airtime.AirtimeService) *orderHandler {
	return &orderHandler{
		store: store,
		airtimeService:airtimeService,
	}
}

type DataStore interface {
	CreateOrder(ctx context.Context, order orders.Order) error
	GetOrder(ctx context.Context, orderId string) (*orders.Order, error)
	GetOrders(ctx context.Context, username string, product string, skipCount int32, maxResultCount int32) ([]*orders.Order, int64, error)
	ChangeStatus(ctx context.Context, orderId string, status orders.OrderStatus) error
	UpdateOrder(ctx context.Context, orderId string, orderDetails map[string]string) error
}

func (o orderHandler) PlaceOrder(ctx context.Context, req *orders.PlaceOrderRequest, resp *orders.PlaceOrderResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("unable to generate uuid, %s", err.Error())
	}

	var orderAttributes []string
	switch req.Product {
	case Airtime:
		airtimeAttributes, err := o.airtimeService.OrderAttributes(ctx, &airtime.EmptyMessage{})
		if err == nil {
			orderAttributes = airtimeAttributes.Attributes
		}
		break
	}

	for _, attribute := range orderAttributes {
		if _, found := req.Details[attribute]; !found {
			return fmt.Errorf("cannot place order, required attribute, %s not sent", strings.Replace(attribute, "_", " ", -1))
		}
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

func (o orderHandler) MarkOrderAsPaid(ctx context.Context, req *orders.MarkOrderAsPaidRequest, resp *orders.EmptyMessage) error {
	order, err := o.store.GetOrder(ctx, req.OrderId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid order id")
		}

		return err
	}

	if order.Amount < req.Amount {
		if err = o.store.ChangeStatus(ctx, order.Id, orders.OrderStatus_UnderPaid); err != nil{
			//TODO: fatal error, notify admin
		}
		return errors.New("amount received is less down the order amount, please contact the admin for resolution")
	}

	switch order.Product {
	case Airtime:
		sendAirtimeReq := &airtime.SendAirtimeRequest{}
		if phoneNumber, found := order.Details["phone_number"]; found {
			sendAirtimeReq.PhoneNumber = phoneNumber
		}

		if network, found := order.Details["provider_id"]; found {
			sendAirtimeReq.ProviderId = network
		}

		sendAirtimeReq.Value = req.Amount

		sendAirtimeResp, err := o.airtimeService.SendAirtime(ctx, sendAirtimeReq)
		if err != nil {
			if changeStatusErr := o.store.ChangeStatus(ctx, order.Id, orders.OrderStatus_ErrorSendingValue); changeStatusErr != nil{
				//TODO: fatal error, notify admin
				log.Error(changeStatusErr)
			}
			// TODO: fatal error, notify admin
			return fmt.Errorf("payment received but unable to fulfil your order, %s", err.Error())
		}

		fmt.Println(sendAirtimeResp)
		order.Details["transaction_id"] = sendAirtimeResp.TransactionId
		break
	}

	if err = o.store.ChangeStatus(ctx, order.Id, orders.OrderStatus_Completed); err != nil{
		//TODO: fatal error, notify admin
	}

	// TODO: order fulfil, notify the customer


	if err = o.store.UpdateOrder(ctx, order.Id, order.Details); err != nil{
		//TODO: fatal error, notify admin
	}
	return nil
}

func (o orderHandler) ChangeStatus(ctx context.Context, req *orders.ChangeStatusRequest, resp *orders.EmptyMessage) error {
	return o.store.ChangeStatus(ctx, req.OrderId, req.Status)
}

func (o orderHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest, resp *orders.GetOrdersResponse) error {
	orderSlice, totalCount, err := o.store.GetOrders(ctx, req.Username, req.Product, req.SkipCount, req.MaxResultCount)
	if err != nil {
		return err
	}

	resp.Orders = orderSlice
	resp.TotalCount = totalCount
	return nil
}

func (o orderHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest, resp *orders.GetOrderResponse) error {
	order, err := o.store.GetOrder(ctx, req.OrderId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid order id")
		}

		return err
	}

	resp.Order = order
	return nil
}

