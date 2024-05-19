package rate_usecase

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"exchange_rate/pkg/packages/errors"
	pstgrs "exchange_rate/pkg/repository/psql"
	rate_dto "exchange_rate/pkg/usecase/rate/dto"
)

func (u *RateUC) CreateRate(
	ctx context.Context,
	data *rate_dto.CreateRateDto,
) *errors.Error {
	return u.rep.CreateRateService(
		ctx,
		rate_entity.NewRate(data.Rate, data.R030, data.ExchangeDate, data.Cc),
	)
}

func (u *RateUC) CreateOrUpdateRate(
	ctx context.Context,
	data *rate_dto.CreateRateDto,
) *errors.Error {
	rate, err := u.rep.GetRateByValCodeService(ctx, data.Cc)
	if err != nil && err.GetMessage() != pstgrs.ErrNotFound(pstgrs.Rate).Message {
		return err
	}

	if rate != nil {
		u.UpdateRate(ctx, rate.Update(data.Rate, data.ExchangeDate))
	}

	return u.CreateRate(ctx, data)
}

func (u *RateUC) GetRateByValCode(
	ctx context.Context,
	data *rate_dto.GetRateByValCodeDto,
) (*rate_entity.Rate, *errors.Error) {
	return u.rep.GetRateByValCodeService(ctx, data.ValCode)
}

func (u *RateUC) GetRates(
	ctx context.Context,
) ([]*rate_entity.Rate, *errors.Error) {
	return u.rep.GetRatesService(ctx)
}

func (u *RateUC) UpdateRate(
	ctx context.Context,
	data *rate_entity.Rate,
) *errors.Error {
	return u.rep.UpdateRateService(ctx, data)
}

func (u *RateUC) DeleteRateByValCode(
	ctx context.Context,
	data *rate_dto.DeleteRateByValCodeDto,
) *errors.Error {
	return u.rep.DeleteRateByValCodeService(ctx, data.ValCode)
}
