package scheduler_service

import (
	"exchange_rate/pkg/packages/errors"
	"exchange_rate/pkg/utils"
	"fmt"
	"time"
)

type Config struct {
	updatedCurrencyMailConfig
	updateRatesConfig
}

type updatedCurrencyMailConfig struct {
	SendUpdatedCurrencyMailAt string
}
type updateRatesConfig struct {
	UpdateRatesTimer string
}

func NewSchedulerConfig() (*Config, *errors.Error) {
	envData := utils.TryGetEnv[string]

	updateRatesTimer, err := envData("UPDATE_RATES_TIMER")
	if err != nil {
		return nil, newErrorNoEnvVar("UPDATE_RATES_TIMER")
	}

	scheduleTime, err := envData("SCHEDULED_TIME")
	if err != nil {
		return nil, newErrorNoEnvVar("SCHEDULED_TIME")
	}

	t, errB := time.Parse("15:04", scheduleTime)
	if errB != nil {
		return nil, errors.NewCriticalErrorf("invalid time format: %v", errB)
	}
	sendMailAt := fmt.Sprintf("%d %d * * *", t.Minute(), t.Hour())

	fmt.Printf("--> CRON EXPR: %s\n", sendMailAt)
	fmt.Printf("--> TIME NOW: %s\n", time.Now())
	return &Config{
		updatedCurrencyMailConfig: updatedCurrencyMailConfig{
			SendUpdatedCurrencyMailAt: sendMailAt,
		},
		updateRatesConfig: updateRatesConfig{
			UpdateRatesTimer: updateRatesTimer,
		},
	}, nil
}
