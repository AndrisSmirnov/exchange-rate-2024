package app

import (
	"context"
	"exchange_rate/pkg/infrastructure/currency"
	"exchange_rate/pkg/infrastructure/currency/nbu"
	"exchange_rate/pkg/infrastructure/mail"
	cron_usecase "exchange_rate/pkg/usecase/cron"
	mail_usecase "exchange_rate/pkg/usecase/mail"
	rate_usecase "exchange_rate/pkg/usecase/rate"
	user_usecase "exchange_rate/pkg/usecase/user"

	"exchange_rate/pkg/packages/errors"

	"github.com/sirupsen/logrus"
)

type UseCases struct {
	*cron_usecase.CronUC
	*mail_usecase.MailUC
	*rate_usecase.RateUC
	*user_usecase.UserUC
}

func NewUseCases(
	ctx context.Context,
	rep *AppRepository,
	basicValCode string,
) (*UseCases, *errors.Error) {
	mailer, err := mail.NewEmailService()
	if err != nil {
		return nil, err
	}

	nbuService, err := nbu.NewServiceNBU()
	if err != nil {
		return nil, err
	}

	mailUC, err := initMailUC(ctx, mailer)
	if err != nil {
		return nil, err
	}

	rateUC, err := initRateUC(ctx, rep)
	if err != nil {
		return nil, err
	}

	userUC, err := initUserUC(ctx, rep)
	if err != nil {
		return nil, err
	}
	cronUC, err := initCronUC(ctx, mailUC, rateUC, userUC, nbuService, basicValCode)
	if err != nil {
		return nil, err
	}

	logrus.Infof("✔ %s", "All use case initialized")

	useCase := &UseCases{
		CronUC: cronUC,
		MailUC: mailUC,
		RateUC: rateUC,
		UserUC: userUC,
	}

	return useCase, nil
}

func initMailUC(
	ctx context.Context,
	sender *mail.EmailSender,
) (
	*mail_usecase.MailUC, *errors.Error,
) {
	infra := mail_usecase.Infrastructure{
		Mail: sender,
	}
	mailUseCase, err := mail_usecase.NewMailUC(
		ctx,
		&infra,
	)

	return mailUseCase, err
}

func initRateUC(
	ctx context.Context,
	repository *AppRepository,
) (
	*rate_usecase.RateUC, *errors.Error,
) {
	db := struct {
		rate_usecase.RateStore
		rate_usecase.RateService
	}{
		RateStore:   repository.Repository,
		RateService: repository.RateService,
	}

	rep := &rate_usecase.Repository{
		Store: db,
	}

	rateUseCase, err := rate_usecase.NewRateUC(
		ctx,
		rep,
	)

	return rateUseCase, err
}

func initUserUC(
	ctx context.Context,
	repository *AppRepository,
) (
	*user_usecase.UserUC, *errors.Error,
) {
	db := struct {
		user_usecase.UserStore
		user_usecase.UserService
	}{
		UserStore:   repository.Repository,
		UserService: repository.UserService,
	}

	rep := &user_usecase.Repository{
		Store: db,
	}

	rateUseCase, err := user_usecase.NewUserUC(
		ctx,
		rep,
	)

	return rateUseCase, err
}

func initCronUC(
	ctx context.Context,
	mailUC *mail_usecase.MailUC,
	rateUC *rate_usecase.RateUC,
	userUC *user_usecase.UserUC,
	currency currency.Currency,
	basicValCode string,
) (
	*cron_usecase.CronUC, *errors.Error,
) {
	services := &cron_usecase.Services{
		MailService: mailUC,
		RateService: rateUC,
		UserService: userUC,
	}

	infra := &cron_usecase.Infrastructure{
		Currency: currency,
	}

	musicVideoUseCase, err := cron_usecase.NewCronUC(
		ctx,
		cron_usecase.NewCronUseCaseConfig(basicValCode),
		services,
		infra,
	)

	return musicVideoUseCase, err
}
