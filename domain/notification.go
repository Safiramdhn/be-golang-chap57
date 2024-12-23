package domain

type Notification struct {
	ID        string `json:"id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	UserEmail string `json:"user_id"`
	OrderID   int    `json:"order_id"`
	PaymentID int    `json:"payment_id"`
	Type      string `json:"type"`
	IsSend    bool   `json:"is_send"`
}

func NotificationData() []Notification {
	return []Notification{
		{
			ID:        "1",
			Subject:   "Order Confirmation",
			Body:      "Your order has been confirmed",
			UserEmail: "user1@example.com",
			OrderID:   1,
			Type:      "order_notification",
			IsSend:    false,
		},
		{
			ID:        "2",
			Subject:   "Payment Confirmation",
			Body:      "Your payment has been confirmed",
			UserEmail: "user2@example.com",
			PaymentID: 1,
			Type:      "payment_notification",
			IsSend:    false,
		},
		{
			ID:        "3",
			Subject:   "Order Confirmation",
			Body:      "Your order has been confirmed",
			UserEmail: "user3@example.com",
			OrderID:   2,
			Type:      "order_notification",
			IsSend:    false,
		},
	}
}
