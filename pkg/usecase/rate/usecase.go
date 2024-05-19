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
		rep: rep,
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
	return nil
}

func (uc *RateUC) checkConfig() *errors.Error {
	return nil
}
