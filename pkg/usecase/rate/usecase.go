package rate_usecase

import (
	"context"
	errors "exchange_rate/pkg/packages/errors"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	Store
}

type Services struct {
}

type Infrastructure struct {
}

type Config struct {
	ServerType, SystemID string
}

type RateUC struct {
	ctx context.Context
	rep *Repository
}

func NewRateUC(
	ctx context.Context,
	rep *Repository,
) (*RateUC, *errors.Error) {
	uc := &RateUC{
		ctx: ctx,
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

	if err := uc.checkConfig(); err != nil {
		return nil, err
	}

	logrus.Infof("✔ %s\n", "Mail use case is initialized")

	return uc, nil
}

func (uc *RateUC) checkStore() *errors.Error {
	if uc.rep == nil {
		return newErrorStoreNotFound("Store!")
	}
	if uc.rep.Store == nil {
		return newErrorStoreNotFound("Store Mail")
	}

	return nil
}

func (uc *RateUC) checkServices() *errors.Error {
	// if uc.services == nil {
	// 	return newErrorServiceNotFound("Services!")
	// }
	// if uc.services.Presenter == nil {
	// 	return newErrorServiceNotFound("Presenter service")
	// }

	return nil
}

func (uc *RateUC) checkConfig() *errors.Error {
	// if uc.conf == nil {
	// 	return newErrorConfig("Config is nil pointer")
	// }
	// if len(uc.conf.SystemID) == 0 {
	// 	return newErrorConfig("SystemID is empty string")
	// }
	// if len(uc.conf.ServerType) == 0 {
	// 	return newErrorConfig("ServerType is empty string")
	// }

	return nil
}
