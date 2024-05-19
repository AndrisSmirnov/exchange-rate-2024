package repository

import (
	"context"
	"time"

	pstgrs "exchange_rate/pkg/repository/psql"
	rate_service "exchange_rate/pkg/repository/service/rate"
	user_service "exchange_rate/pkg/repository/service/user"
	"exchange_rate/pkg/utils"
)

type Repository struct {
	RateService *rate_service.RateService
	UserService *user_service.UserService
	PstgrsUser  *pstgrs.Repository
}

type Config struct {
	psqlConfig *pstgrs.Config
}

func NewConfig() (*Config, error) {
	envData := utils.TryGetEnv[string]

	psqlUser, err := envData("POSTGRES_USER")
	if err != nil {
		return nil, err
	}

	psqlPassword, err := envData("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}

	psqlNetwork, err := envData("POSTGRES_NETWORK")
	if err != nil {
		return nil, err
	}

	psqlHost, err := envData("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}

	psqlPort, err := envData("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}

	psqlDriver, err := envData("DRIVER_NAME")
	if err != nil {
		return nil, err
	}

	psqlSSL, err := envData("SSL_MODE")
	if err != nil {
		return nil, err
	}

	dbName, err := envData("POSTGRES_PORT_DB_NAME")
	if err != nil {
		return nil, err
	}

	psqlMaxCons := utils.TryGetEnvDefault[int]("POSTGRES_MAX_IDLE_CONS", 1000)

	psqlLifeTime := utils.TryGetEnvDefault[int]("POSTGRES_MAX_IDLE_CON_LIFETIME_MINUTES", 1)

	c := &Config{
		psqlConfig: pstgrs.NewConfig(
			psqlUser,
			psqlPassword,
			psqlNetwork,
			psqlHost,
			psqlPort,
			dbName,
			psqlSSL,
			psqlMaxCons,
			time.Minute,
			time.Duration(psqlLifeTime)*time.Minute,
			psqlDriver,
		),
	}

	return c, nil
}

func New(
	ctx context.Context,
	errChan chan error,
	conf Config,
) (*Repository, error) {
	psql, err := pstgrs.NewRepository(ctx, conf.psqlConfig, errChan)
	if err != nil {
		return nil, err
	}

	return &Repository{
		RateService: rate_service.NewRateService(psql),
		UserService: user_service.NewUserService(psql),
		PstgrsUser:  psql,
	}, nil
}
