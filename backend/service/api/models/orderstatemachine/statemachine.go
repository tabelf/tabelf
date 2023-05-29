package orderstatemachine

import (
	"context"
	"tabelf/backend/service/api/pkg/statemachine"
	"tabelf/backend/service/app"
)

func HandleEvent(ctx context.Context, state string, event string, args map[string]interface{}) (
	err error,
) {
	fsm := TakeawayFSM(state)
	if fsm.Cannot(event) {
		return app.ErrCustomerOrderEventCanNot(ctx)
	}
	err = fsm.Event(ctx, event, args)
	if err != nil {
		return err
	}
	return nil
}

func TakeawayFSM(initialState string) *statemachine.FSM {
	return statemachine.NewFSM(
		initialState,
		statemachine.Events{
			{
				Name:     OrderPaymentSuccessEvent, // 订单支付成功, 待支付到已支付
				Src:      OrderUnpaidState,
				Dst:      OrderPaidState,
				Callback: PayFailed,
			},
			{
				Name:     OrderPaymentFailEvent, // 订单支付成功, 待支付到已支付
				Src:      OrderUnpaidState,
				Dst:      OrderPayFailState,
				Callback: PaySuccess,
			},
			{
				Name:     OrderUnpaidTimeoutEvent, // 订单支付取消, 待支付到订单取消
				Src:      OrderUnpaidState,
				Dst:      OrderCancelState,
				Callback: Cancel,
			},
		},
	)
}
