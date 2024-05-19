package pstgrs

import (
	"exchange_rate/pkg/packages/errors"
	"fmt"
	"net/http"

	"github.com/lib/pq"
)

var ErrDuplicateKeyValue = pq.ErrorCode("23505")

var (
	ErrInternalUnknown = errors.NewInternalError(
		"psql: unknown error in operation",
	).AddCode(http.StatusInternalServerError)
	ErrRequest = errors.NewInternalError(
		"internal server error in processing request to DB",
	).AddCode(http.StatusInternalServerError)
)

func ErrNotFound(name string) *errors.Error {
	return errors.New(
		fmt.Sprintf("%s not found", name)).AddCode(http.StatusNotFound)
}

func ErrAlreadyExists(name string) *errors.Error {
	return errors.New(
		fmt.Sprintf("%s already exists", name)).AddCode(http.StatusConflict)
}
