package pstgrs

import (
	"context"
	"fmt"
	"time"

	"exchange_rate/pkg/packages/errors"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	TRY_ATTEMPTS                    = 3
	SECONDS_ADD_ON_FAIL             = 60
	HEALTH_CHECK_PSQL_PING_INTERVAL = 30
)

type Postgres struct {
	DB      *sqlx.DB
	Config  *Config
	errChan chan error
}

func newConnection(
	ctx context.Context,
	config *Config,
	errChan chan error,
) (*Postgres, *errors.Error) {
	conn, err := New(ctx, config, errChan)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func New(ctx context.Context, cfg *Config, errChan chan error) (*Postgres, *errors.Error) {
	conn, err := connectToPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}
	pg := &Postgres{
		DB:      conn,
		Config:  cfg,
		errChan: errChan,
	}
	pg.startHealthCheck(ctx)
	return pg, nil
}

type PostgresQueryer struct {
	Postgres
	Queryer *sqlx.DB
}

func NewQueryer(ctx context.Context, cfg *Config, errChan chan error) (*PostgresQueryer, error) {
	pg, err := New(ctx, cfg, errChan)
	if err != nil {
		return nil, err
	}
	q := &PostgresQueryer{
		Postgres: *pg,
		Queryer:  pg.DB,
	}
	return q, nil
}

func (p *Postgres) reconnect(ctx context.Context) error {
	conn, err := connectToPostgres(ctx, p.Config)
	if err != nil {
		return err
	}
	p.DB = conn
	return nil
}

func connectToPostgres(_ context.Context, cfg *Config) (*sqlx.DB, *errors.Error) {
	logrus.Info("Connecting to Postgres...")

	var (
		db  *sqlx.DB
		err error
	)

	for i := 1; i <= TRY_ATTEMPTS; i++ {
		t := SECONDS_ADD_ON_FAIL * i
		db, err = openConnection(cfg)
		if err == nil {
			break
		}

		logrus.Warnf(
			"Error connect to %s. Try %d. Next try after %d seconds.",
			cfg.DriverName, i, t,
		)

		time.Sleep(time.Duration(t) * time.Second)
	}

	if err != nil {
		return nil, errors.NewCriticalError(
			"psql: error connect to Postgres",
		).AddError(err)
	}
	if db == nil {
		return nil, errors.NewCriticalError(
			"psql: error connect to Postgres. db is nil",
		)
	}

	logrus.Info("Connected to Postgres")

	return db, nil
}

func openConnection(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DriverName, buildDSN(cfg))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleCons)
	db.SetConnMaxIdleTime(cfg.MaxIdleConnDuration)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DSN - data source name
func buildDSN(cfg *Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func (c *Postgres) startHealthCheck(ctx context.Context) {
	ticker := time.NewTicker(
		time.Duration(HEALTH_CHECK_PSQL_PING_INTERVAL) * time.Second,
	)

	go func() {
		for {
			<-ticker.C
			if err := c.DB.Ping(); err != nil {
				c.errChan <- errors.NewInternalError("error ping Postgres").AddError(err)
				logrus.Info("Reconnecting to PostgresDB...")
				if err := c.reconnect(ctx); err != nil {
					c.errChan <- errors.NewCriticalError("error reconnect to PostgresDB").AddError(err)
				}
			}
		}
	}()
}
