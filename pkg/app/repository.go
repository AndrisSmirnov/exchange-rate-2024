package app

import (
	"context"
	"exchange_rate/pkg/repository"
	pstgrs "exchange_rate/pkg/repository/psql"
	rate_service "exchange_rate/pkg/repository/service/rate"
	user_service "exchange_rate/pkg/repository/service/user"
)

type AppRepository struct {
	*rate_service.RateService
	*user_service.UserService
	*pstgrs.Repository
}

func newRepository(
	ctx context.Context, errChan chan error,
) (
	*AppRepository, error,
) {
	config, err := repository.NewConfig()
	if err != nil {
		return nil, err
	}

	rep, err := repository.New(ctx, errChan, *config)
	if err != nil {
		return nil, err
	}

	return &AppRepository{
		rep.RateService,
		rep.UserService,
		rep.PstgrsUser,
	}, nil
}
