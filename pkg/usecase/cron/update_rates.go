package cron_usecase

import (
	rate_entity "exchange_rate/pkg/domain/rate/entity"
	"time"

	"github.com/sirupsen/logrus"
)

func (u *CronUC) UpdateAllRates() {
	rates, err := u.services.RateService.GetRates(u.ctx)
	if err != nil {
		logrus.Error("error getting rates in cron job. restarting in 1 minute")
		go u.restartUpdateAllRates()
		return
	}

	toUpdateMap := make(map[string]*rate_entity.Rate, len(rates))
	for _, r := range rates {
		toUpdateMap[r.ValCode] = r
	}

	reqRates := make([]string, 0, len(rates))
	for _, r := range rates {
		reqRates = append(reqRates, r.ValCode)
	}

	updatedRatesMap := u.infra.Currency.GetUpdatedCurrency(reqRates)
	for valCode, updateRataData := range updatedRatesMap {
		if rate, ok := toUpdateMap[valCode]; ok {
			rate = rate.Update(updateRataData.Rate, updateRataData.ExchangeDate)
			if err := u.services.RateService.UpdateRate(u.ctx, rate); err != nil {
				logrus.Errorf("Error updating rate with ID: %s. Reason: %s", valCode, err)
			}
		}
	}
}

func (u *CronUC) restartUpdateAllRates() {
	time.Sleep(1 * time.Minute)
	go u.SendToAllUsersCurrencyUpdateForUsdUah()
}
