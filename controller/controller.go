package controller

import (
	service "be-chap56/services"

	"go.uber.org/zap"
)

type Controller struct {
	Notification NotificationController
	Log          *zap.Logger
}

func NewController(service service.Service, log *zap.Logger) *Controller {
	return &Controller{
		Notification: *NewNotificationController(service.Email, log),
	}
}
