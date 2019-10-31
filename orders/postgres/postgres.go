package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ademuanthony/bitenvoy/orders/postgres/models"
	"github.com/ademuanthony/bitenvoy/orders/proto/orders"
	"github.com/ademuanthony/bitenvoy/utils"
	"github.com/gofrs/uuid"
	"github.com/micro/go-micro/util/log"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func NewPgDb(host, port, user, pass, dbname string) (*PgDb, error) {
	db, err := utils.PgConnect(host, port, user, pass, dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	return &PgDb{
		db: db,
		queryTimeout: time.Second * 30,
	}, nil
}

type PgDb struct {
	db *sql.DB
	queryTimeout time.Duration
}

func (pg *PgDb) CreateOrder(ctx context.Context, order go_micro_srv_orders.Order) error {
	details, err := json.Marshal(order.Details)
	if err != nil {
		return fmt.Errorf("cannot place order, error in converting order details to json string, %s", err.Error())
	}

	orderModel := models.OrderModel{
		ID:       order.Id,
		Username: order.Username,
		Product:  order.Product,
		Status:   order.Status.String(),
		Details:  string(details),
		Date:     int(order.Date),
		Amount:   order.Amount,
	}

	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("cannot place order, error in starting transaction, %s", err.Error())
	}

	rollback := func() {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Errorf("order not created and error occurred in tx rollback, %s", rollBackErr.Error())
		}
	}

	if err := orderModel.Insert(ctx, tx, boil.Infer()); err != nil {
		return fmt.Errorf("cannot place order, %s", err.Error())
	}

	id, err := uuid.NewV4()
	if err != nil {
		rollback()
		return fmt.Errorf("cannot place order, error in generating status if, %s", err.Error())
	}

	orderStatus := models.OrderStatus{
		ID:      id.String(),
		OrderID: order.Id,
		Status:  order.Status.String(),
		Date:    int(time.Now().UTC().Unix()),
	}

	if err := orderStatus.Insert(ctx, tx, boil.Infer()); err != nil {
		rollback()
		return fmt.Errorf("cannot place order, error in saving order status, %s", err.Error())
	}

	return nil
}

func (pg *PgDb) GetOrder(ctx context.Context, orderId string) (*go_micro_srv_orders.Order, error) {
	orderModel, err := models.OrderModels(models.OrderModelWhere.ID.EQ(orderId)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	var orderDetails map[string]string
	if err = json.Unmarshal([]byte(orderModel.Details), &orderDetails); err != nil {
		return nil, fmt.Errorf("cannot get order details from json, %s", err.Error())
	}

	order := go_micro_srv_orders.Order{
		Id:       orderModel.ID,
		Username: orderModel.Username,
		Product:  orderModel.Product,
		Amount:   orderModel.Amount,
		Date:     int64(orderModel.Date),
		Details:  orderDetails,
		Status:   go_micro_srv_orders.OrderStatus(go_micro_srv_orders.OrderStatus_value[orderModel.Status]),
	}

	return &order, nil
}

func (pg *PgDb) GetOrders(ctx context.Context, username string, product string, skipCount int32, maxResultCount int32) ([]*go_micro_srv_orders.Order, int64, error) {
	var query []qm.QueryMod
	if username != "" {
		query = append(query, models.OrderModelWhere.Username.EQ(username))
	}

	if product != "" {
		query = append(query, models.OrderModelWhere.Product.EQ(product))
	}

	total, err := models.OrderModels(query...).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	query = append(query, qm.Offset(int(skipCount)), qm.Limit(int(maxResultCount)))

	orderSlice, err := models.OrderModels(query...).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var result []*go_micro_srv_orders.Order
	for _, order := range orderSlice {
		var orderDetails map[string]string
		if err = json.Unmarshal([]byte(order.Details), &orderDetails); err != nil {
			return nil, 0, fmt.Errorf("cannot get order details from json, %s", err.Error())
		}

		result = append(result, &go_micro_srv_orders.Order{
			Id:                   order.ID,
			Username:             order.Username,
			Product:              order.Product,
			Amount:               order.Amount,
			Date:                 int64(order.Date),
			Details:              orderDetails,
			Status:   go_micro_srv_orders.OrderStatus(go_micro_srv_orders.OrderStatus_value[order.Status]),
		})
	}

	return result, total, nil
}

func (pg *PgDb) ChangeStatus(ctx context.Context, orderId string, status go_micro_srv_orders.OrderStatus) error {
	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("cannot change status, error in starting transaction, %s", err.Error())
	}

	rollback := func() {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Errorf("order not created and error occurred in tx rollback, %s", rollBackErr.Error())
		}
	}

	orderModel, err := models.OrderModels(models.OrderModelWhere.ID.EQ(orderId)).One(ctx, tx)
	if err != nil {
		return err
	}

	orderModel.Status = status.String()
	_, err = orderModel.Update(ctx, tx, boil.Infer())

	if err != nil {
		return err
	}

	id, err := uuid.NewV4()
	if err != nil {
		rollback()
		return fmt.Errorf("error in generating order status id, %s", err.Error())
	}

	orderStatus := models.OrderStatus{
		ID:      id.String(),
		OrderID: orderId,
		Status:  status.String(),
		Date:    int(time.Now().UTC().Unix()),
	}

	if err = orderStatus.Insert(ctx, tx, boil.Infer()); err != nil {
		rollback()
		return fmt.Errorf("error in changing order status, %s", err.Error())
	}

	err = tx.Commit()
	return err
}

func (pg *PgDb) UpdateOrder(ctx context.Context, orderId string, orderDetails map[string]string) error  {
	orderModel, err := models.OrderModels(models.OrderModelWhere.ID.EQ(orderId)).One(ctx, pg.db)
	if err != nil {
		return err
	}

	details, err := json.Marshal(orderDetails)
	if err != nil {
		return fmt.Errorf("error in converting order details to json string, %s", err.Error())
	}

	orderModel.Details = string(details)
	_, err = orderModel.Update(ctx, pg.db, boil.Infer())
	return err
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}
