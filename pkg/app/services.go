package app

import (
	"exchange_rate/pkg/packages/errors"
	scheduler_service "exchange_rate/pkg/services/scheduler"
)

type Services struct {
	SchedulerService scheduler_service.SchedulerService
}

func newServices(uc *UseCases) (*Services, *errors.Error) {
	schedulerServiceConfig, err := scheduler_service.NewSchedulerConfig()
	if err != nil {
		return nil, nil
	}

	return &Services{
		SchedulerService: scheduler_service.NewSchedulerService(schedulerServiceConfig, uc),
	}, nil
}
