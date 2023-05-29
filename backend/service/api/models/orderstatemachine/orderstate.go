package orderstatemachine

const (
	OrderUnpaidState  = "Unpaid"   // 待支付.
	OrderPaidState    = "Paid"     // 已支付.
	OrderPayFailState = "PayFail"  // 支付失败.
	OrderCancelState  = "Canceled" // 订单已取消.
)
