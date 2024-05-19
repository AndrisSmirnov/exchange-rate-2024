package cron_usecase

func NewCronUseCaseConfig(basicValCode string) *Config {
	return &Config{
		BasicValCode: basicValCode,
	}
}
