package mail_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/domain/vo"
	"exchange_rate/pkg/packages/errors"
)

type (
	Sender interface {
		SendEmail(_ context.Context, data *rate_entity.Rate, receivers ...vo.Email) *errors.Error
	}
)
