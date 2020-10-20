package gowaves

type OrderbookCancelRequest struct {
	Sender    string `json:"sender"`
	OrderID   string `json:"orderId"`
	Timestamp int    `json:"timestamp"`
	Signature string `json:"signature"`
}
