package mapper

import (
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	rate_model "exchange_rate/pkg/domain/rate/model"
	"exchange_rate/pkg/domain/vo"
)

func ConvertRateToDB(rate *rate_entity.Rate) *rate_model.RateDB {
	return &rate_model.RateDB{
		ID:      rate.ID.String(),
		Date:    rate.Date,
		Rate:    rate.Rate,
		ValCode: rate.ValCode,
		Code:    rate.Code,
	}
}

func ConvertRateFromDB(rateDB *rate_model.RateDB) *rate_entity.Rate {
	return &rate_entity.Rate{
		ID:      vo.UUID(rateDB.ID),
		Date:    rateDB.Date,
		Rate:    rateDB.Rate,
		ValCode: rateDB.ValCode,
		Code:    rateDB.Code,
	}
}
