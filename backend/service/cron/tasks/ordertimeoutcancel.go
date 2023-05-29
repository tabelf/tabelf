package tasks

import (
	"log"
	entpayorder "tabelf/backend/gen/entschema/payorder"
	"tabelf/backend/service/api/models/orderstatemachine"
	"tabelf/backend/service/app"
	"time"
)

func OrderTimeoutCancel(jobCtx JobContext, config app.Config) {
	current := time.Now()
	ctx := jobCtx.Context
	// 订单超过 15 分钟未支付，自动取消
	orders, err := app.EntClient.PayOrder.Query().Where(
		entpayorder.DeactivatedAtIsNil(),
		entpayorder.Status(orderstatemachine.OrderUnpaidState),
		entpayorder.CreatedAtGTE(current.Add(-8*time.Hour)),
		entpayorder.CreatedAtLTE(current.Add(-app.OrderExpiredAt)),
	).All(ctx)
	if err != nil {
		return
	}
	if len(orders) == 0 {
		log.Printf("not find unpaid order.")
		return
	}
	log.Printf("start order timeout cancel function. orders = %v", orders)
	for _, order := range orders {
		if err = orderstatemachine.HandleEvent(ctx,
			order.Status,
			orderstatemachine.OrderUnpaidTimeoutEvent,
			map[string]interface{}{app.UserOrderKey: order},
		); err != nil {
			continue
		}
	}
	log.Printf("end order timeout cancel function.")
}
