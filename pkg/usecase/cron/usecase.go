package cron_usecase

import (
	"context"
	"exchange_rate/pkg/packages/errors"

	"github.com/sirupsen/logrus"
)

type Repository struct {
}

type Services struct {
	MailService MailService
	RateService RateService
	UserService UserService
}

type Infrastructure struct {
	Currency CurrencyInfrastructure
}

type Config struct {
	ServerType, SystemID string
	BasicValCode         string
}

type CronUC struct {
	ctx      context.Context
	conf     *Config
	services *Services
	infra    *Infrastructure
}

func NewCronUC(
	ctx context.Context,
	conf *Config,
	services *Services,
	infra *Infrastructure,
) (*CronUC, *errors.Error) {
	uc := &CronUC{
		ctx:      ctx,
		conf:     conf,
		services: services,
		infra:    infra,
	}

	if uc.ctx == nil {
		return nil, newErrorContext()
	}

	if err := uc.checkStore(); err != nil {
		return nil, err
	}

	if err := uc.checkServices(); err != nil {
		return nil, err
	}

	if err := uc.checkInfrastructures(); err != nil {
		return nil, err
	}

	if err := uc.checkConfig(); err != nil {
		return nil, err
	}

	logrus.Infof("✔ %s\n", "Cron use case is initialized")

	return uc, nil
}

func (uc *CronUC) checkStore() *errors.Error {
	return nil
}

func (uc *CronUC) checkServices() *errors.Error {
	if uc.services == nil {
		return newErrorServiceNotFound("Services!")
	}
	if uc.services.MailService == nil {
		return newErrorServiceNotFound("Mail service")
	}
	if uc.services.RateService == nil {
		return newErrorServiceNotFound("Rate service")
	}
	if uc.services.UserService == nil {
		return newErrorServiceNotFound("User service")
	}

	return nil
}

func (uc *CronUC) checkInfrastructures() *errors.Error {
	if uc.infra == nil {
		return newErrorServiceNotFound("Infrastructure!")
	}
	if uc.infra.Currency == nil {
		return newErrorServiceNotFound("Currency infrastructure")
	}

	return nil
}

func (uc *CronUC) checkConfig() *errors.Error {
	if uc.conf == nil {
		return newErrorConfig("Config is nil pointer")
	}
	if len(uc.conf.BasicValCode) == 0 {
		return newErrorConfig("BasicValCode is empty string")
	}

	return nil
}
