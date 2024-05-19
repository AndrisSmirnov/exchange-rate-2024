package rate_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	rate_model "exchange_rate/pkg/domain/rate/model"
	errors "exchange_rate/pkg/packages/errors"
)

type (
	Store interface {
		RateStore
		RateService
	}

	RateStore interface {
		CreateRate(ctx context.Context, data *rate_model.RateDB) *errors.Error
		GetRateByID(ctx context.Context, id string) (*rate_model.RateDB, *errors.Error)
		GetRateByValCode(ctx context.Context, valCode string) (*rate_model.RateDB, *errors.Error)
		GetRates(ctx context.Context) ([]*rate_model.RateDB, *errors.Error)
		UpdateRate(ctx context.Context, data *rate_model.RateDB) *errors.Error
		DeleteRateByID(ctx context.Context, id string) *errors.Error
		DeleteRateByValCode(ctx context.Context, vadCode string) *errors.Error
	}

	RateService interface {
		CreateRateService(ctx context.Context, rate *rate_entity.Rate) *errors.Error
		GetRateByIdService(ctx context.Context, id string) (*rate_entity.Rate, *errors.Error)
		GetRateByValCodeService(ctx context.Context, valCode string) (*rate_entity.Rate, *errors.Error)
		GetRatesService(ctx context.Context) ([]*rate_entity.Rate, *errors.Error)
		UpdateRateService(ctx context.Context, rate *rate_entity.Rate) *errors.Error
		DeleteRateByIdService(ctx context.Context, id string) *errors.Error
		DeleteRateByValCodeService(ctx context.Context, valCode string) *errors.Error
	}
)
