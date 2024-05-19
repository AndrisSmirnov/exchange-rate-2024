package http

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/packages/errors"
	rate_dto "exchange_rate/pkg/usecase/rate/dto"
	user_dto "exchange_rate/pkg/usecase/user/dto"
)

type (
	Services interface {
		Rate
		User
	}
	Rate interface {
		GetRateByValCode(
			ctx context.Context,
			data *rate_dto.GetRateByValCodeDto,
		) (*rate_entity.Rate, *errors.Error)
	}
	User interface {
		CreateUser(ctx context.Context, data *user_dto.CreateUserDto) *errors.Error
	}
)
