package pstgrs

import (
	"context"
	"exchange_rate/pkg/packages/errors"
)

const (
	User = "user"
	Rate = "rate"
)

type Repository struct {
	Postgres
}

func NewRepository(
	ctx context.Context,
	config *Config,
	errChan chan error,
) (*Repository, *errors.Error) {
	conn, err := newConnection(ctx, config, errChan)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Postgres: *conn,
	}, nil
}
