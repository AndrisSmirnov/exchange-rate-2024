package user_usecase

import (
	"context"
	"exchange_rate/pkg/packages/errors"

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

type UserUC struct {
	ctx   context.Context
	rep   *Repository
	infra Infrastructure
}

func NewUserUC(
	ctx context.Context,
	rep *Repository,
) (*UserUC, *errors.Error) {
	uc := &UserUC{
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

	logrus.Infof("✔ %s\n", "User use case is initialized")

	return uc, nil
}

func (uc *UserUC) checkStore() *errors.Error {
	if uc.rep == nil {
		return newErrorStoreNotFound("Store!")
	}
	if uc.rep.Store == nil {
		return newErrorStoreNotFound("Store User")
	}

	return nil
}

func (uc *UserUC) checkServices() *errors.Error {
	return nil
}

func (uc *UserUC) checkConfig() *errors.Error {
	return nil
}
