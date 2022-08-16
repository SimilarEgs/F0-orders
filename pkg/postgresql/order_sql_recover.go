package postgresql

import (
	"fmt"

	"github.com/SimilarEgs/L0-orders/config"
	"github.com/SimilarEgs/L0-orders/internal/models"
	"github.com/SimilarEgs/L0-orders/pkg/cache"
	"github.com/SimilarEgs/L0-orders/pkg/constants"
)

func (db *DB) Recover(cfg *config.Config) error {

	db.Init(cfg)

	order := models.Order{}
	orderDelivey := models.Delivery{}
	orderPayment := models.Payment{}
	orderItems := make([]models.Items, 0)

	orderQuery := fmt.Sprintf(`SELECT * FROM %s`, ordersTable)

	deliveryQuery := fmt.Sprintf(`
	SELECT name, phone, zip, city, address, region, email
	FROM %s
	WHERE order_uid = $1`, deliveryTable)

	paymentQuery := fmt.Sprintf(`
	SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
	FROM %s
	WHERE order_uid = $1`, paymentTable)

	itemsQuery := fmt.Sprintf(`
	SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
	FROM %s
	WHERE order_uid = $1`, itemsTable)

	orderRes, err := db.Con.Query(orderQuery)
	if err != nil {
		return err
	}
	defer orderRes.Close()

	for orderRes.Next() {

		orderRes.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.ShardKey, &order.SmID, &order.DateCreated, &order.OOFShard)

		deliveryRes, err := db.Con.Query(deliveryQuery, order.OrderUID)
		if err != nil {
			return err
		}
		defer deliveryRes.Close()

		for deliveryRes.Next() {

			deliveryRes.Scan(&orderDelivey.Name, &orderDelivey.Phone, &orderDelivey.Zip, &orderDelivey.City, &orderDelivey.Address, &orderDelivey.Region,
				&orderDelivey.Email)
		}
		order.Delivery = orderDelivey

		paymentRes, err := db.Con.Query(paymentQuery, order.OrderUID)
		if err != nil {
			return err
		}
		defer paymentRes.Close()

		for paymentRes.Next() {

			paymentRes.Scan(&orderPayment.Transaction, &orderPayment.RequestID, &orderPayment.Currency,
				&orderPayment.Provider, &orderPayment.Amount, &orderPayment.PaymentDT, &orderPayment.Bank,
				&orderPayment.DeliveryCost, &orderPayment.GoodsTotal, &orderPayment.CustomFee)
		}

		order.Payment = orderPayment

		itemsRes, err := db.Con.Query(itemsQuery, order.OrderUID)
		if err != nil {
			return err
		}
		defer paymentRes.Close()

		for itemsRes.Next() {

			item := models.Items{}

			itemsRes.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name,
				&item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)

			orderItems = append(orderItems, item)
		}
		order.Items = orderItems

		cache.AppCache.Set(order.OrderUID, order, constants.CacheDuration)
	}

	return nil
}
