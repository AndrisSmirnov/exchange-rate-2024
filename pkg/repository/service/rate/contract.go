package rate_service

import (
	"context"
	rate_model "exchange_rate/pkg/domain/rate/model"
	"exchange_rate/pkg/packages/errors"
)

type (
	Repository interface {
		CreateRate(ctx context.Context, data *rate_model.RateDB) *errors.Error
		GetRateByID(ctx context.Context, id string) (*rate_model.RateDB, *errors.Error)
		GetRateByValCode(ctx context.Context, valCode string) (*rate_model.RateDB, *errors.Error)
		GetRates(ctx context.Context) ([]*rate_model.RateDB, *errors.Error)
		UpdateRate(ctx context.Context, data *rate_model.RateDB) *errors.Error
		DeleteRateByID(ctx context.Context, id string) *errors.Error
		DeleteRateByValCode(ctx context.Context, valCode string) *errors.Error
	}
)
