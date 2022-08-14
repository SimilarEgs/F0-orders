package models

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type Order struct {
	Order_uid          string `json:"order_uid"           validate:"required"`
	Track_number       string `json:"track-number"`
	Entry              string `json:"entry"`
	Locate             string `json:"locate"`
	Internal_signature string `json:"internal_signature"`
	Customer_id        string `json:"customer_id"`
	Delivery_service   string `json:"delivery_service"`
	Shardkey           string `json:"shardkey"`
	Sm_id              int    `json:"sm_id"               validate:"required"`
	Date_created       string `json:"date_created"`
	OOF_shard          string `json:"oof_shard"`

	Delivery Delivery
	Payment  Payment
	Items    []Items
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"        validate:"numeric"`
	PaymentDT    int    `json:"payment_dt"    validate:"numeric"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost" validate:"numeric"`
	GoodsTotal   int    `json:"goods_total"   validate:"numeric"`
	CustomFee    int    `json:"custom_fee"    validate:"numeric"`
}

type Items struct {
	ChrtID      int    `json:"chrt_id"       validate:"numeric"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"         validate:"numeric"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"          validate:"numeric"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"   validate:"numeric"`
	NmID        int    `json:"nm_id"         validate:"numeric"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"        validate:"numeric"`
}

func (o Order) ValidateOrder() error {

	validate = validator.New()

	return validate.Struct(o)
}
