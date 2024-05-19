package mail_usecase

import (
	"context"
	"exchange_rate/pkg/packages/errors"

	"github.com/sirupsen/logrus"
)

type Repository struct {
}

type Services struct {
}

type Infrastructure struct {
	Mail Sender
}

type Config struct {
	ServerType, SystemID string
}

type MailUC struct {
	ctx   context.Context
	infra *Infrastructure
}

func NewMailUC(
	ctx context.Context,
	infra *Infrastructure,
) (*MailUC, *errors.Error) {
	uc := &MailUC{
		ctx:   ctx,
		infra: infra,
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

func (uc *MailUC) checkStore() *errors.Error {
	return nil
}

func (uc *MailUC) checkServices() *errors.Error {
	return nil
}

func (uc *MailUC) checkInfrastructures() *errors.Error {
	if uc.infra == nil {
		return newErrorServiceNotFound("Infrastructure!")
	}
	if uc.infra.Mail == nil {
		return newErrorServiceNotFound("Mail infrastructure")
	}

	return nil
}

func (uc *MailUC) checkConfig() *errors.Error {
	return nil
}
