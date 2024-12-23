package controller

import (
	"be-chap56/domain"
	service "be-chap56/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type NotificationController struct {
	EmailService service.EmailService
	Log          *zap.Logger
}

func NewNotificationController(service service.EmailService, log *zap.Logger) *NotificationController {
	return &NotificationController{
		EmailService: service,
		Log:          log,
	}
}

func (ctrl *NotificationController) SendOrderMail(c *gin.Context) {
	var orderNotif domain.Notification
	if err := c.BindJSON(&orderNotif); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	orderNotif.Type = "order_notification"
	orderNotif.ID = uuid.New().String()
	orderNotif.Subject = "Order Confirmation"
	orderNotif.IsSend = false
	_, err := ctrl.EmailService.Send(orderNotif.UserEmail, orderNotif.Subject, "order", orderNotif)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	orderNotif.IsSend = true
	c.JSON(200, gin.H{"message": "Order notification sent successfully"})
}

func (ctrl *NotificationController) SendPaymentMail(c *gin.Context) {
	var paymentNotif domain.Notification
	if err := c.BindJSON(&paymentNotif); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paymentNotif.Type = "payment_notification"
	paymentNotif.ID = uuid.New().String()
	paymentNotif.Subject = "Payment Confirmation"
	paymentNotif.IsSend = false
	_, err := ctrl.EmailService.Send(paymentNotif.UserEmail, paymentNotif.Subject, "payment", paymentNotif)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	paymentNotif.IsSend = true
	c.JSON(200, gin.H{"message": "Payment notification sent successfully"})
}

func (ctrl *NotificationController) All(c *gin.Context) {
	notifications := domain.NotificationData()
	c.JSON(200, gin.H{"data": notifications})
}
