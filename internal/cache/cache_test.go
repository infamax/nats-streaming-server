package cache

import (
	"github.com/infamax/nats-streaming-server/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	cache := New()
	order := models.Order{
		OrderUid:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: models.Delivery(struct {
			Name    string
			Phone   string
			Zip     string
			City    string
			Address string
			Region  string
			Email   string
		}{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		}),
		Payment: models.Payment(struct {
			Transaction  string
			RequestID    string
			Currency     string
			Provider     string
			Amount       int
			PaymentDt    int
			Bank         string
			DeliveryCost float64
			GoodsTotal   int
			CustomFee    float64
		}{
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
		}),
		Items: []models.Item{
			models.Item{
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
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       time.Time{},
		OofShard:          "1",
	}

	go func() {
		cache.Add(&order)
	}()

	time.Sleep(1 * time.Second)

	orderTest, err := cache.Get("b563feb7b2b84b6test")
	assert.Nil(t, err)
	assert.Equal(t, order, *orderTest)
}
