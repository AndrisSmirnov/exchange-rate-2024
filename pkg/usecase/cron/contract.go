package cron_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	user_entity "exchange_rate/pkg/domain/user/entity"
	"exchange_rate/pkg/domain/vo"
	rate_dto "exchange_rate/pkg/usecase/rate/dto"

	"exchange_rate/pkg/packages/errors"
)

type (
	MailService interface {
		SendUpdatedCurrencyEmail(
			ctx context.Context,
			rate *rate_entity.Rate,
			emails []vo.Email,
		) *errors.Error
	}
	RateService interface {
		CreateOrUpdateRate(ctx context.Context, data *rate_dto.CreateRateDto) *errors.Error
		GetRateByValCode(
			ctx context.Context, data *rate_dto.GetRateByValCodeDto) (*rate_entity.Rate, *errors.Error)
		GetRates(ctx context.Context) ([]*rate_entity.Rate, *errors.Error)
		UpdateRate(ctx context.Context, data *rate_entity.Rate) *errors.Error
	}
	UserService interface {
		GetUsers(ctx context.Context) ([]*user_entity.User, *errors.Error)
	}
	CurrencyInfrastructure interface {
		GetUpdatedCurrency([]string) map[string]rate_dto.CreateRateDto
	}
)
