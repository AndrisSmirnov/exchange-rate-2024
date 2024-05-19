package app

import (
	"context"

	"exchange_rate/pkg/controllers"
	"exchange_rate/pkg/packages/errors"

	"exchange_rate/pkg/utils"
)

type App struct {
	rep         *AppRepository
	useCases    *UseCases
	controllers *controllers.Controllers
	services    *Services

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
	serverURL, err := utils.TryGetEnv[string]("SERVER_URL")
	if err != nil {
		return nil, errors.New("empty SERVER_URL")
	}

	useCases, err := newUseCases(ctx, rep, basicValCode)
	if err != nil {
		return nil, err
	}

	controllers, err := createControllers(useCases, serverURL, basicValCode)
	if err != nil {
		return nil, err
	}

	services, err := newServices(useCases)
	if err != nil {
		return nil, err
	}

	app := &App{
		rep:         rep,
		useCases:    useCases,
		controllers: controllers,
		services:    services,

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
	app.controllers.Start()
	app.services.SchedulerService.StartAllJob()
	return nil
}

func (app *App) Stop() {
	if app.listener != nil {
		app.listener.Stop()
	}
	app.services.SchedulerService.StopAllJob()
}
