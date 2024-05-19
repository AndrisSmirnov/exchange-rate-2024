package currency

import (
	"exchange_rate/pkg/infrastructure/currency/nbu"
	rate_dto "exchange_rate/pkg/usecase/rate/dto"
)

type currency struct {
	*nbu.NBU
}

type Currency interface {
	GetUpdatedCurrency([]string) map[string]rate_dto.CreateRateDto
}

func NewCurrency(
	nbuCur *nbu.NBU,
) Currency {
	return currency{NBU: nbuCur}
}
