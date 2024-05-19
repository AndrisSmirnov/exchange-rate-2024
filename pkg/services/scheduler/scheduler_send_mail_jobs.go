package scheduler_service

import (
	"fmt"
)

func (ss *schedulerService) StartMailJob() {
	ss.schedulers[mail].Start()
}

func (ss *schedulerService) StopMailJob() {
	ss.schedulers[mail].Stop()
}

// UpdatedCurrencyRates microservice
func (ss *schedulerService) setupMailJobs() {
	id, err := ss.schedulers[mail].AddFunc(
		fmt.Sprintf("@every %ss", ss.config.updatedCurrencyMailConfig.SendUpdatedCurrencyMailTimer),
		ss.schedulerService.SendToAllUsersCurrencyUpdateForUsdUah,
	)
	handleSchedulerError(err)

	jobNames[id] = "Send Updated Currency Rates To Mail"
}
