package repository

import (
	"context"
	"github.com/infamax/nats-streaming-server/internal/models"
	"log"
)

func (d *db) GetByUUID(ctx context.Context, uuid string) (*models.Order, error) {
	log.Println("getting data from table")
	const query = `select
		id, order_uid, track_number, entry,
		locale, internal_signature, customer_id,
    	delivery_service, shardkey, sm_id,
		date_created, oof_shard
		from orders
		where order_uid = $1;`
	var order models.Order
	var id int
	log.Println("query start")
	err := d.pool.QueryRow(ctx, query, uuid).Scan(&id,
		&order.OrderUid, &order.TrackNumber, &order.Entry,
		&order.Locale, &order.InternalSignature, &order.CustomerId,
		&order.DeliveryService, &order.Shardkey, &order.SmId,
		&order.DateCreated, &order.OofShard)
	log.Println("query finish")
	log.Println("id = ", id)
	log.Println("uuid = ", order.OrderUid)
	log.Println("trackNumber = ", order.TrackNumber)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	items, err := d.getItems(ctx, order.TrackNumber)
	log.Println("err = ", err)
	if err != nil {
		return nil, err
	}
	delivery, err := d.getDelivery(ctx, id)
	log.Println("id = ", id)
	if err != nil {
		return nil, err
	}
	payment, err := d.getPayment(ctx, uuid)
	if err != nil {
		log.Println("err = ", err)
		log.Println("error!")
		return nil, err
	}
	log.Println("err = ", err)
	order.Items = items
	order.Delivery = *delivery
	order.Payment = *payment
	return &order, nil
}

func (d *db) GetByID(ctx context.Context, id int) (string, error) {
	const query = `
		select data
		from invalid_messages
		where id = $1;
	`
	var data string
	err := d.pool.QueryRow(ctx, query, id).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (d *db) getItems(ctx context.Context, trackNumber string) ([]models.Item, error) {
	const query = `
		select chrt_id, track_number,
		price, rid, name, sale, size,
		total_price, nm_id, brand, status
		from items
		where track_number = $1;
	`

	rows, err := d.pool.Query(ctx, query, trackNumber)
	if err != nil {
		return nil, err
	}
	var items []models.Item
	for rows.Next() {
		var item models.Item
		_ = rows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price,
			&item.Rid, &item.Name, &item.Sale, &item.Size,
			&item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		items = append(items, item)
	}
	return items, nil
}

func (d *db) getPayment(ctx context.Context, uuid string) (*models.Payment, error) {
	const query = `
		select transaction, request_id, currency,
        provider, amount, payment_dt, bank, delivery_coast,
    	goods_total, custom_fee
		from payments
		where transaction = $1;
	`
	var payment models.Payment
	err := d.pool.QueryRow(ctx, query, uuid).Scan(
		&payment.Transaction, &payment.RequestID, &payment.Currency,
		&payment.Provider, &payment.Amount, &payment.PaymentDt,
		&payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal,
		&payment.CustomFee)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (d *db) getDelivery(ctx context.Context, id int) (*models.Delivery, error) {
	const query = `
		select name, phone,
		zip, city, address, region,
    	email
		from delivery
		where order_id = $1;
	`
	var delivery models.Delivery
	err := d.pool.QueryRow(ctx, query, id).Scan(&delivery.Name,
		&delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address,
		&delivery.Region, &delivery.Email)
	if err != nil {
		return nil, err
	}
	return &delivery, nil
}
