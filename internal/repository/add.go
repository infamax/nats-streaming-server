package repository

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
)

func (d *db) AddModel(ctx context.Context, order *models.Order) error {
	const query = `
		insert into orders (order_uid, track_number, entry,
		locale, internal_signature, customer_id, delivery_service,
    	shardkey, sm_id, date_created, oof_shard)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		returning id;
	`
	var id int
	err := d.pool.QueryRow(ctx, query, order.OrderUid, order.TrackNumber, order.Entry,
		order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService,
		order.Shardkey, order.SmId, order.DateCreated, order.OofShard).Scan(&id)
	if err != nil {
		return err
	}
	err = d.addItems(ctx, order.Items)
	if err != nil {
		return err
	}
	err = d.addDelivery(ctx, id, &order.Delivery)
	if err != nil {
		return err
	}
	err = d.addPayment(ctx, &order.Payment)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) AddData(ctx context.Context, data string) (int, error) {
	const query = `insert into 
			invalid_messages (data)
			values($1)
			returning id;
	`
	var id int
	err := d.pool.QueryRow(ctx, query, data).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *db) addDelivery(ctx context.Context, id int, delivery *models.Delivery) error {
	const query = `
		insert into delivery
		(order_id, name, phone, zip,
   		 city, address, region, email)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		returning order_id;
	`

	err := d.pool.QueryRow(ctx, query, id, delivery.Name, delivery.Phone,
		delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

func (d *db) addPayment(ctx context.Context, payment *models.Payment) error {
	const query = `
		insert into payments (transaction, request_id, currency,
   		 provider, amount, payment_dt, bank, delivery_coast, goods_total,
    	custom_fee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning transaction;
	`

	err := d.pool.QueryRow(ctx, query, payment.Transaction, payment.RequestID,
		payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt,
		payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee).Scan(&payment.Transaction)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) addItems(ctx context.Context, items []models.Item) error {
	const query = `
		insert into items (chrt_id, track_number, price,
		rid, name, sale, size, total_price, nm_id, brand,
    	status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		returning chrt_id;
	`

	for _, item := range items {
		err := d.pool.QueryRow(ctx, query, item.ChrtId, item.TrackNumber, item.Price,
			item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId,
			item.Brand, item.Status).Scan(&item.ChrtId)
		if err != nil {
			return err
		}
	}
	return nil
}
