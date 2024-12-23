package infra

import (
	"be-chap56/config"
	"be-chap56/controller"
	"be-chap56/log"
	service "be-chap56/services"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cfg config.Config
	Ctl controller.Controller
	Log *zap.Logger
}

func NewServiceContext() (*ServiceContext, error) {
	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig()
	if err != nil {
		return handlerError(err)
	}

	// instance logger
	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		return handlerError(err)
	}

	// instance service
	services := service.NewService(appConfig, logger)

	// instance controller
	Ctl := controller.NewController(services, logger)

	return &ServiceContext{Cfg: appConfig, Ctl: *Ctl, Log: logger}, nil
}
