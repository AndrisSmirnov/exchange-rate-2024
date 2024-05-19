package rate_service

import (
	"context"
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	rate_mapper "exchange_rate/pkg/domain/rate/mapper"
	"exchange_rate/pkg/packages/errors"
)

func (s *RateService) CreateRateService(
	ctx context.Context, rate *rate_entity.Rate,
) *errors.Error {
	return s.repository.CreateRate(ctx, rate_mapper.ConvertRateToDB(rate))
}

func (s *RateService) GetRateByIdService(
	ctx context.Context, id string,
) (
	*rate_entity.Rate, *errors.Error,
) {
	rateDB, err := s.repository.GetRateByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return rate_mapper.ConvertRateFromDB(rateDB), nil
}

func (s *RateService) GetRateByValCodeService(
	ctx context.Context, valCode string,
) (
	*rate_entity.Rate, *errors.Error,
) {
	rateDB, err := s.repository.GetRateByValCode(ctx, valCode)
	if err != nil {
		return nil, err
	}

	return rate_mapper.ConvertRateFromDB(rateDB), nil
}

func (s *RateService) GetRatesService(
	ctx context.Context,
) (
	[]*rate_entity.Rate, *errors.Error,
) {
	ratesDB, err := s.repository.GetRates(ctx)
	if err != nil {
		return nil, err
	}

	rates := make([]*rate_entity.Rate, 0, len(ratesDB))

	for _, el := range ratesDB {
		rates = append(rates, rate_mapper.ConvertRateFromDB(el))
	}

	return rates, nil
}

func (s *RateService) UpdateRateService(
	ctx context.Context, rate *rate_entity.Rate,
) *errors.Error {
	return s.repository.UpdateRate(ctx, rate_mapper.ConvertRateToDB(rate))
}

func (s *RateService) DeleteRateByIdService(ctx context.Context, id string) *errors.Error {
	return s.repository.DeleteRateByID(ctx, id)
}

func (s *RateService) DeleteRateByValCodeService(ctx context.Context, valCode string) *errors.Error {
	return s.repository.DeleteRateByValCode(ctx, valCode)
}
