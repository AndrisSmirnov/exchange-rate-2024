package cron_usecase

import (
	"exchange_rate/pkg/domain/vo"
	rate_dto "exchange_rate/pkg/usecase/rate/dto"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func (u *CronUC) SendToAllUsersCurrencyUpdateForUsdUah() {
	data := &rate_dto.GetRateByValCodeDto{
		RateValCode: rate_dto.RateValCode{
			ValCode: u.conf.BasicValCode,
		},
	}
	rate, err := u.services.RateService.GetRateByValCode(u.ctx, data)
	if err != nil {
		logrus.Error("error getting rate in cron job. restarting in 1 minute")
		go u.restartToSendToAllUsersCurrencyUpdateForUsdUah()
		return
	}

	users, err := u.services.UserService.GetUsers(u.ctx)
	if err != nil {
		logrus.Error("error getting users in cron job. restarting in 1 minute")
		go u.restartToSendToAllUsersCurrencyUpdateForUsdUah()
		return
	}

	emails := make([]vo.Email, 0, len(users))
	for _, u := range users {
		emails = append(emails, u.Mail)
	}

	fmt.Printf("--> RATE: %v\n", rate)
	fmt.Printf("--> EMAILS: %v\n", emails)

	if err := u.services.MailService.SendUpdatedCurrencyEmail(u.ctx, rate, emails); err != nil {
		logrus.Errorf("error sending currency to users. Reason: %v", err)
	}
}

func (u *CronUC) restartToSendToAllUsersCurrencyUpdateForUsdUah() {
	time.Sleep(1 * time.Minute)
	go u.SendToAllUsersCurrencyUpdateForUsdUah()
}
