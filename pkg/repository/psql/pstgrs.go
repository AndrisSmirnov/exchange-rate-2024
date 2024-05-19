package pstgrs

import (
	"context"
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
) (*Repository, error) {
	conn, err := newConnection(ctx, config, errChan)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Postgres: *conn,
	}, nil
}
