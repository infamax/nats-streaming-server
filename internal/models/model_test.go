package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validJson = `{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}`

func TestParseValidJson(t *testing.T) {
	var order Order
	err := json.Unmarshal([]byte(validJson), &order)
	if err != nil {
		t.Error("json must be valid")
	}

	assert.Equal(t, "b563feb7b2b84b6test", order.OrderUid)
	assert.Equal(t, "WBILMTESTTRACK", order.TrackNumber)
	assert.Equal(t, "WBIL", order.Entry)
	assert.Equal(t, "en", order.Locale)
	assert.Equal(t, "", order.InternalSignature)
	assert.Equal(t, "test", order.CustomerId)
	assert.Equal(t, "meest", order.DeliveryService)
	assert.Equal(t, "9", order.Shardkey)
	assert.Equal(t, 99, order.SmId)
	dateCreated, err := order.DateCreated.MarshalText()
	assert.Nil(t, err)
	assert.Equal(t, "2021-11-26T06:22:19Z", string(dateCreated))
	assert.Equal(t, "1", order.OofShard)

	item := Item{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}

	assert.Equal(t, item, order.Items[0])

	delivery := Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}

	assert.Equal(t, delivery, order.Delivery)

	payment := Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}

	assert.Equal(t, payment, order.Payment)
}

var invalidJson = `{
	"name": "anton",
	"school": "32",
}`

func TestParseInvalidJson(t *testing.T) {
	var order Order
	err := json.Unmarshal([]byte(invalidJson), &order)
	if err == nil {
		t.Error("json is not valid")
	}
}
