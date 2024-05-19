package scheduler_service

type Config struct {
	updatedCurrencyMailConfig
	updateRatesConfig
}

type updatedCurrencyMailConfig struct {
	SendUpdatedCurrencyMailTimer string
}
type updateRatesConfig struct {
	UpdateRatesTimer string
}

func NewSchedulerConfig(
	sendUpdatedCurrencyMailTimer, updateRatesTimer string,
) (*Config, error) {
	// envData := utils.TryGetEnv[string]

	// checkPendingGenerationImagesTimer, err := envData("SEND_UPDATED_CURRENCY_MAIL_TIMER")
	// if err != nil {
	// 	return nil, newErrorNoEnvVar("SEND_UPDATED_CURRENCY_MAIL_TIMER")
	// }

	// updateRatesTimer, err := envData("UPDATE_RATES_TIMER")
	// if err != nil {
	// 	return nil, newErrorNoEnvVar("UPDATE_RATES_TIMER")
	// }

	return &Config{
		updatedCurrencyMailConfig: updatedCurrencyMailConfig{
			SendUpdatedCurrencyMailTimer: sendUpdatedCurrencyMailTimer,
		},
		updateRatesConfig: updateRatesConfig{
			UpdateRatesTimer: updateRatesTimer,
		},
	}, nil
}
