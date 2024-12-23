package service

import (
	"be-chap56/config"

	"go.uber.org/zap"
)

type Service struct {
	Email EmailService
}

func NewService(appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Email: NewEmailService(appConfig.Email, log),
	}
}
