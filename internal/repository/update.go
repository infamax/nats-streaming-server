package repository

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
)

func (d *db) UpdateModel(ctx context.Context, order *models.Order) error {
	const query = `
		update orders
		set track_number = $2,
		entry = $3,
    	locale = $4,
    	internal_signature = $5,
    	customer_id = $6,
    	delivery_service = $7,
    	shardkey = $8,
    	sm_id = $9,
    	date_created = $10,
    	oof_shard = $11
		where order_uid = $1;
	`
	_, err := d.pool.Exec(ctx, query, order.OrderUid, order.TrackNumber,
		order.Entry, order.Locale, order.InternalSignature, order.CustomerId,
		order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated,
		order.OofShard)

	if err != nil {
		return err
	}

	err = d.updateItems(ctx, order)

	if err != nil {
		return err
	}
	err = d.updatePayment(ctx, order)

	if err != nil {
		return err
	}

	err = d.updateDelivery(ctx, order)

	if err != nil {
		return err
	}

	return nil
}

func (d *db) UpdateData(ctx context.Context, id int, data string) error {
	const query = `
		update invalid_messages
		set data = $2
		where id = $1;
	`
	_, err := d.pool.Exec(ctx, query, id, data)
	return err
}

func (d *db) updateItems(ctx context.Context, order *models.Order) error {
	const query = `
		update items
		set chrt_id = $2,
    	price = $3,
    	rid = $4,
    	name = $5,
    	sale = $6,
    	size = $7,
    	total_price = $8,
    	nm_id = $9,
    	brand = $10,
    	status = $11
		where track_number = $1;
	`

	for _, item := range order.Items {
		_, err := d.pool.Exec(ctx, query, item.TrackNumber,
			item.ChrtId, item.Price, item.Rid, item.Name,
			item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *db) updatePayment(ctx context.Context, order *models.Order) error {
	const query = `
		update payments
		set request_id = $2,
    	currency = $3,
    	provider = $4,
    	amount = $5,
    	payment_dt = $6,
    	bank = $7,
    	delivery_coast = $8,
    	goods_total = $9,
    	custom_fee = $10
		where transaction = $1;
	`

	_, err := d.pool.Exec(ctx, query, order.Payment.Transaction, order.Payment.RequestID,
		order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost,
		order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}
	return nil
}

func (d *db) updateDelivery(ctx context.Context, order *models.Order) error {
	const querySelectId = `
		select id
		from orders
		where order_uid = $1;
	`

	var id int
	err := d.pool.QueryRow(ctx, querySelectId, order.OrderUid).Scan(&id)
	if err != nil {
		return err
	}

	const query = `
		update delivery
		set name = $2,
    	phone = $3,
    	zip = $4,
    	city = $5,
    	address = $6,
    	region = $7,
    	email = $8
		where order_id = $1;
	`

	_, err = d.pool.Exec(ctx, query, id, order.Delivery.Name,
		order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)

	if err != nil {
		return err
	}
	return nil
}
