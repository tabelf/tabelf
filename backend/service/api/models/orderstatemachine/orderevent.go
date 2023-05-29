package orderstatemachine

const (
	OrderPaymentSuccessEvent = "PaymentSuccess" // 订单支付完成.
	OrderPaymentFailEvent    = "PaymentFail"    // 订单支付失败.
	OrderUnpaidTimeoutEvent  = "UnpaidTimeout"  // 待支付超时.
)
