package service

import (
	"context"

	"be-chap57/notification_service/proto"

	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationService struct {
	proto.UnimplementedNotificationServiceServer
}

type Notification struct {
	ID        string `json:"id"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	UserEmail string `json:"user_email"`
	OrderID   int    `json:"order_id,omitempty"`
	PaymentID int    `json:"payment_id,omitempty"`
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

func (s *NotificationService) SendOrderMail(ctx context.Context, req *proto.OrderNotificationRequest) (*proto.NotificationResponse, error) {
	if req.UserEmail == "" {
		return nil, status.Errorf(codes.InvalidArgument, "UserEmail is required")
	}

	notificationID := uuid.New().String()
	log.Printf("Sending order mail to %s with NotificationID: %s", req.UserEmail, notificationID)

	response := &proto.NotificationResponse{
		Notification: &proto.Notification{
			Id:        notificationID,
			UserEmail: req.UserEmail,
			OrderId:   req.OrderId,
			Type:      "order_notification",
			Subject:   "Order Confirmation",
			IsSent:    true,
		},
	}

	return response, nil
}

func (s *NotificationService) SendPaymentMail(ctx context.Context, req *proto.PaymentNotificationRequest) (*proto.NotificationResponse, error) {
	if req.UserEmail == "" {
		return nil, status.Errorf(codes.InvalidArgument, "UserEmail is required")
	}

	notificationID := uuid.New().String()
	log.Printf("Sending payment mail to %s with NotificationID: %s", req.UserEmail, notificationID)

	response := &proto.NotificationResponse{
		Notification: &proto.Notification{
			Id:        notificationID,
			UserEmail: req.UserEmail,
			PaymentId: req.PaymentId,
			Type:      "payment_notification",
			Subject:   "Payment Confirmation",
			IsSent:    true,
		},
	}

	return response, nil
}

func (s *NotificationService) GetAllNotifications(ctx context.Context, req *proto.Empty) (*proto.NotificationListResponse, error) {
	notifications := []*proto.Notification{}
	for _, n := range NotificationData() {
		notifications = append(notifications, &proto.Notification{
			Id:        n.ID,
			UserEmail: n.UserEmail,
			OrderId:   int64(n.OrderID),
			PaymentId: int64(n.PaymentID),
			Type:      n.Type,
			Subject:   n.Subject,
			IsSent:    n.IsSend,
		})
	}

	response := &proto.NotificationListResponse{
		Notifications: notifications,
	}

	return response, nil
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar) {
	proto.RegisterNotificationServiceServer(s, &NotificationService{})
}
