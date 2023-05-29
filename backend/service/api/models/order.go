package models

type OrderPaymentResponse struct {
	QrCodeImage string `json:"qr_code_image"` // 二维码支付地址
}
