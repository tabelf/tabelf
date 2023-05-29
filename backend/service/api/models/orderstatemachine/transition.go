package orderstatemachine

import (
	"context"

	"tabelf/backend/gen/entschema"
	entpayorder "tabelf/backend/gen/entschema/payorder"
	"tabelf/backend/service/api/pkg/statemachine"
	"tabelf/backend/service/app"
)

func Cancel(ctx context.Context, event *statemachine.Event) error {
	anOrder, ok := event.Args[app.UserOrderKey].(*entschema.PayOrder)
	if !ok {
		return app.ErrCustomerOrderInfoNotExist(ctx)
	}
	option := UpdateOrderOption{
		OldStatus:   event.Src,
		Status:      event.Dst,
		CancelEvent: event.Event,
	}
	if err := UpdateOrder(ctx, anOrder, &option); err != nil {
		return err
	}
	return nil
}

func PayFailed(ctx context.Context, event *statemachine.Event) error {
	anOrder, ok := event.Args[app.UserOrderKey].(*entschema.PayOrder)
	if !ok {
		return app.ErrCustomerOrderInfoNotExist(ctx)
	}
	return UpdateOrder(ctx, anOrder, &UpdateOrderOption{
		OldStatus: event.Src,
		Status:    event.Dst,
	})
}

func PaySuccess(ctx context.Context, event *statemachine.Event) error {
	anOrder, ok := event.Args[app.UserOrderKey].(*entschema.PayOrder)
	if !ok {
		return app.ErrCustomerOrderInfoNotExist(ctx)
	}
	transactionNumber, ok := event.Args[app.TransactionNumberKey].(string)
	if !ok {
		return app.ErrCustomerOrderTransactionNumberEmpty(ctx)
	}
	option := UpdateOrderOption{
		OldStatus:         event.Src,
		Status:            event.Dst,
		TransactionNumber: transactionNumber,
	}
	return UpdateOrder(ctx, anOrder, &option)
}

type UpdateOrderOption struct {
	OldStatus   string
	Status      string
	CancelEvent string

	TransactionNumber string // 交易订单号
}

func UpdateOrder(ctx context.Context, anOrder *entschema.PayOrder, option *UpdateOrderOption) (err error) {
	// 如果订单状态等于当前状态不做改变
	if option.Status == anOrder.Status {
		return nil
	}
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		// 乐观锁.
		orderUpdate := tx.PayOrder.Update()
		if app.IsNotBlank(option.Status) {
			orderUpdate.SetStatus(option.Status)
		}
		if app.IsNotBlank(option.TransactionNumber) {
			orderUpdate.SetTransactionNumber(option.TransactionNumber)
		}
		if app.IsNotBlank(option.CancelEvent) {
			orderUpdate.SetCancelEvent(option.CancelEvent)
		}
		if err = orderUpdate.Where(
			entpayorder.ID(anOrder.ID),
			entpayorder.Status(option.OldStatus),
		).Exec(ctx); err != nil {
			app.Log.Error(ctx, err)
			return err
		}
		return nil
	}); err != nil {
		app.Log.Error(ctx, err)
		return err
	}
	return nil
}
