package app

import (
	"context"

	"exchange_rate/pkg/controllers"
	"exchange_rate/pkg/packages/errors"
	user_dto "exchange_rate/pkg/usecase/rate/dto"

	"exchange_rate/pkg/utils"
)

type App struct {
	rep         *AppRepository
	useCases    *UseCases
	controllers *controllers.Controllers

	listener   *errors.Listener
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func New(cancelFunc func(), ctx context.Context) (*App, error) {
	l := errors.NewListener(ctx, cancelFunc)
	errChan := l.GetErrChan()

	rep, err := newRepository(ctx, errChan)
	if err != nil {
		return nil, err
	}

	basicValCode, err := utils.TryGetEnv[string]("BASIC_VALCODE")
	if err != nil {
		return nil, errors.New("empty BASIC_VALCODE")
	}

	useCases, err := NewUseCases(ctx, rep, basicValCode)
	if err != nil {
		return nil, err
	}

	controllers, err := createControllers(useCases, basicValCode)
	if err != nil {
		return nil, err
	}

	app := &App{
		rep:         rep,
		useCases:    useCases,
		controllers: controllers,

		listener:   l,
		cancelFunc: cancelFunc,
		ctx:        ctx,
	}

	return app, nil
}

func (app *App) Run() *errors.Error {
	err := app.listener.Start()
	if err != nil {
		return err
	}
	app.useCases.RateUC.CreateOrUpdateRate(app.ctx, &user_dto.CreateRateDto{
		Cc: "USD",
	})
	return nil
}

func (app *App) Stop() {
	if app.listener != nil {
		app.listener.Stop()
	}
}
