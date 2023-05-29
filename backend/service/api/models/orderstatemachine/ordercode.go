package orderstatemachine

const (
	OrderUnpaidStateCode  = "001" // 待支付.
	OrderCancelStateCode  = "040" // 订单取消
	OrderPayFailStateCode = "050" // 支付失败.
	OrderPaidStateCode    = "100" // 已支付.
)

var StateCodeMap = map[string]string{
	OrderUnpaidState:  OrderUnpaidStateCode,
	OrderCancelState:  OrderCancelStateCode,
	OrderPayFailState: OrderPayFailStateCode,
	OrderPaidState:    OrderPaidStateCode,
}

func StatusToCode(state string) string {
	return StateCodeMap[state]
}
