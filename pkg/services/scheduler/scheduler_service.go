package scheduler_service

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type (
	SchedulerService interface {
		StartAllJob()
		StopAllJob()

		MailJobs
		RatesJobs
	}

	MailJobs interface {
		StartMailJob()
		StopMailJob()
	}
	RatesJobs interface {
		StartRateJob()
		StopRateJob()
	}

	Scheduler interface {
		Mail
		Rates
	}
	Mail interface {
		SendToAllUsersCurrencyUpdateForUsdUah()
	}
	Rates interface {
		UpdateAllRates()
	}

	schedulersPool map[string]*cron.Cron
)

const (
	mail = "mail"
	rate = "rate"
)

type schedulerService struct {
	config           *Config
	schedulers       schedulersPool
	schedulerService Scheduler
}

func NewSchedulerService(
	config *Config,
	scheduler Scheduler,
) SchedulerService {
	service := &schedulerService{
		config:           config,
		schedulers:       make(map[string]*cron.Cron),
		schedulerService: scheduler,
	}

	service.schedulers[mail] = cron.New(
		cron.WithLogger(newLogger()),
	)
	service.schedulers[rate] = cron.New(
		cron.WithLogger(newLogger()),
	)

	service.setupMailJobs()
	service.setupRateJobs()

	return service
}

func (ss *schedulerService) StartAllJob() {
	ss.StartMailJob()
	ss.StartRateJob()
}

func (ss *schedulerService) StopAllJob() {
	ss.StopMailJob()
	ss.StopRateJob()
}

func handleSchedulerError(err error) {
	if err != nil {
		logrus.Errorf("Error occurred on add func to scheduler: %v", err)
	}
}
