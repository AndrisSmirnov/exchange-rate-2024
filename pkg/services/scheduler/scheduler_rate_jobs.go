package scheduler_service

import (
	"fmt"
)

func (ss *schedulerService) StartRateJob() {
	ss.schedulers[rate].Start()
}

func (ss *schedulerService) StopRateJob() {
	ss.schedulers[rate].Stop()
}

func (ss *schedulerService) setupRateJobs() {
	id, err := ss.schedulers[rate].AddFunc(
		fmt.Sprintf("@every %ss", ss.config.updateRatesConfig.UpdateRatesTimer),
		ss.schedulerService.UpdateAllRates,
	)
	handleSchedulerError(err)

	jobNames[id] = "Update Rates Process"
}
