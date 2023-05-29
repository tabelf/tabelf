package utils

func GetSessionKey(openID string) string {
	return "session_key:" + openID
}

func GetDraftGrouponKey() string {
	return "draft_groupon"
}

func GetDraftGrouponGlobalIDKey() string {
	return "draft_groupon_global_id"
}

func GetGrouponPurchaseKey(orderNumber string) string {
	return "groupon_purchase:" + orderNumber
}

func GetGrouponRefundKey(orderNumber string) string {
	return "groupon_refund:" + orderNumber
}

func GetGrouponPurchaseByteDanceKey(orderNumber string) string {
	return "groupon_purchase_bytedance:" + orderNumber
}

func GetGrouponRefundByteDanceKey(orderNumber string) string {
	return "groupon_refund_bytedance:" + orderNumber
}
