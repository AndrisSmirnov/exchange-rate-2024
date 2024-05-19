package rate_entity

import "exchange_rate/pkg/domain/vo"

type Rate struct {
	ID      vo.UUID `json:"id"`
	Date    string  `json:"date"`
	Rate    float64 `json:"rate"`
	ValCode string  `json:"valCode"`
	Code    int     `json:"code"`
}

func NewRate(
	rate float64, code int,
	date, valCode string,
) *Rate {
	return &Rate{
		ID:      vo.NewID(),
		Date:    date,
		Rate:    rate,
		ValCode: valCode,
		Code:    code,
	}
}

func (r *Rate) Update(rate float64, date string) *Rate {
	r.Rate = rate
	r.Date = date
	return r
}
