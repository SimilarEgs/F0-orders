package postgresql

import (
	"fmt"

	"github.com/SimilarEgs/L0-orders/internal/models"
)

func (db *DB) Insert(order *models.Order) {

	orderQuery := fmt.Sprintf(`
	INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, ordersTable)

	deliveryQuery := fmt.Sprintf(`
	INSERT INTO %s (name, phone, zip, city, address, region, email, order_uid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, deliveryTable)

	paymentQuery := fmt.Sprintf(`
	INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, paymentTable)

	itemsQuery := fmt.Sprintf(`
	INSERT INTO %s (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, itemsTable)

	db.Con.QueryRow(orderQuery, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SmID, order.DateCreated, order.OOFShard)

	db.Con.QueryRow(deliveryQuery, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region,
		order.Delivery.Email, order.OrderUID)

	db.Con.QueryRow(paymentQuery, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)

	for _, item := range order.Items {
		db.Con.QueryRow(itemsQuery, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID,
			item.Brand, item.Status, order.OrderUID)
	}

}
