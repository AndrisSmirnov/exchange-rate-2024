package pstgrs

import (
	"database/sql"
	"exchange_rate/pkg/packages/errors"
	"net"

	"github.com/lib/pq"
)

func (pg *Postgres) ErrorHandler(err error, name string) *errors.Error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return ErrNotFound(name)
	}
	if internalErr, ok := err.(*pq.Error); ok {
		return pg.handlingPgError(internalErr, name)
	}
	if netOpErr, ok := err.(*net.OpError); ok {
		return pg.handlingNetOpError(netOpErr)
	}
	pg.errChan <- ErrInternalUnknown.AddError(err)
	return ErrRequest.AddError(err)
}

func (pg *Postgres) handlingPgError(err *pq.Error, name string) *errors.Error {
	c := err.Code.Class()
	switch c {
	case pq.ErrorClass("00"):
		return nil
	case pq.ErrorClass("22"):
		pg.errChan <- errors.NewInternalError(
			"psql: invalid data in Postgres request",
		).AddError(err)
	case pq.ErrorClass("08"):
		pg.errChan <- errors.NewCriticalError(
			"psql: disconnected from Postgres",
		).AddError(err)
	case pq.ErrorClass("23"):
		return pg.handlingPgViolationError(err, name)
	default:
		pg.errChan <- ErrInternalUnknown.AddError(err)
	}
	return ErrRequest.AddError(err)
}

func (pg *Postgres) handlingPgViolationError(err *pq.Error, name string) *errors.Error {
	c := err.Code
	switch c {
	case ErrDuplicateKeyValue:
		return ErrAlreadyExists(name)
	default:
		pg.errChan <- ErrInternalUnknown.AddError(err)
		return ErrRequest.AddError(err)
	}
}

func (pg *Postgres) handlingNetOpError(err *net.OpError) *errors.Error {
	if err.Temporary() || err.Timeout() {
		pg.errChan <- errors.NewInternalError(
			"psql: temporary connections problem",
		).AddError(err)
	} else {
		pg.errChan <- errors.NewCriticalError(
			"psql: connections problem",
		).AddError(err)
	}
	return ErrRequest.AddError(err)
}
