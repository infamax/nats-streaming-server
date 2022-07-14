package models

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestID    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       int     `json:"amount"`
	PaymentDt    int     `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}
