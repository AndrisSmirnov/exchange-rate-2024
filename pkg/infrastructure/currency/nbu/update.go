package nbu

import (
	"exchange_rate/pkg/utils"

	rate_dto "exchange_rate/pkg/usecase/rate/dto"

	"github.com/sirupsen/logrus"
)

func (s *NBU) GetUpdatedCurrency(symbols []string) map[string]rate_dto.CreateRateDto {
	resp := map[string]rate_dto.CreateRateDto{}

	for _, symbol := range symbols {
		url := *s.cnf.BaseLink
		q := url.Query()
		q.Set("valcode", symbol)
		q.Set("json", "")
		url.RawQuery = q.Encode()
		res, err := utils.SendRequest[[]rate_dto.CreateRateDto]("GET", url.String(), []byte{})

		if err != nil {
			logrus.Errorf("Failed to get exchange rate for symbol %s. Reason: %v", symbol, err)
		}
		if res == nil || len(*res) == 0 {
			logrus.Errorf("Failed to get exchange rate for symbol %s. Reason: empty response", symbol)
		}

		elements := *res
		resp[symbol] = elements[0]
	}

	return resp
}
