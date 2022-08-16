package models

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type Order struct {
	OrderUID          string `json:"order_uid"           validate:"required"`
	TrackNumber       string `json:"track_number"        validate:"required"`
	Entry             string `json:"entry"               validate:"required"`
	Locale            string `json:"locale"              validate:"required"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"         validate:"required"`
	DeliveryService   string `json:"delivery_service"    validate:"required"`
	ShardKey          string `json:"shardkey"            validate:"required"`
	SmID              int    `json:"sm_id"               validate:"required"`
	DateCreated       string `json:"date_created"        validate:"required"`
	OOFShard          string `json:"oof_shard"           validate:"required"`

	Delivery Delivery `json:"delivery" validate:"required"`
	Payment  Payment  `json:"payment"  validate:"required"`
	Items    []Items  `json:"items"    validate:"required"`
}

type Delivery struct {
	Name    string `json:"name"    validate:"required"`
	Phone   string `json:"phone"   validate:"required"`
	Zip     string `json:"zip"     validate:"required"`
	City    string `json:"city"    validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region"  validate:"required"`
	Email   string `json:"email"   validate:"required"`
}

type Payment struct {
	Transaction  string `json:"transaction"   validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"      validate:"required"`
	Provider     string `json:"provider"      validate:"required"`
	Amount       int    `json:"amount"        validate:"required,numeric"`
	PaymentDT    int    `json:"payment_dt"    validate:"required,numeric"`
	Bank         string `json:"bank"          validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"required,numeric"`
	GoodsTotal   int    `json:"goods_total"   validate:"required,numeric"`
	CustomFee    int    `json:"custom_fee"    validate:"numeric"`
}

type Items struct {
	ChrtID      int    `json:"chrt_id"       validate:"required,numeric"`
	TrackNumber string `json:"track_number"  validate:"required"`
	Price       int    `json:"price"         validate:"required,numeric"`
	Rid         string `json:"rid"           validate:"required"`
	Name        string `json:"name"          validate:"required"`
	Sale        int    `json:"sale"          validate:"required,numeric"`
	Size        string `json:"size"          validate:"required"`
	TotalPrice  int    `json:"total_price"   validate:"required,numeric"`
	NmID        int    `json:"nm_id"         validate:"required,numeric"`
	Brand       string `json:"brand"         validate:"required"`
	Status      int    `json:"status"        validate:"required,numeric"`
}

func (o Order) ValidateOrder() error {

	validate = validator.New()

	return validate.Struct(o)
}
