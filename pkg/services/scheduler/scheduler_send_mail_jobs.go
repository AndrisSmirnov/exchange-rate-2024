package scheduler_service

func (ss *schedulerService) StartMailJob() {
	ss.schedulers[mail].Start()
}

func (ss *schedulerService) StopMailJob() {
	ss.schedulers[mail].Stop()
}

func (ss *schedulerService) setupMailJobs() {
	id, err := ss.schedulers[mail].AddFunc(
		ss.config.updatedCurrencyMailConfig.SendUpdatedCurrencyMailAt,
		ss.schedulerService.SendToAllUsersCurrencyUpdateForUsdUah,
	)
	handleSchedulerError(err)

	jobNames[id] = "Send Updated Currency Rates To Mail"
}
