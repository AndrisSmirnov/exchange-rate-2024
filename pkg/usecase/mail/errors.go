package mail_usecase

import (
	errors "exchange_rate/pkg/packages/errors"
	"net/http"
)

var (
	errAccessDenied = errors.New("access denied").AddCode(http.StatusForbidden)
)

func newErrorContext() *errors.Error {
	return errors.NewCriticalErrorf(
		"init Mail use case. ctx in nil.",
	).AddCode(http.StatusInternalServerError)
}

func newErrorStoreNotFound(name string) *errors.Error {
	return errors.NewCriticalErrorf(
		"init Mail use case. %s is nil pointer.",
		name,
	).AddCode(http.StatusInternalServerError)
}

func newErrorServiceNotFound(name string) *errors.Error {
	return errors.NewCriticalErrorf(
		"init Mail use case. %s is nil pointer.",
		name,
	).AddCode(http.StatusInternalServerError)
}

func newErrorInfrastructureNotFound(name string) *errors.Error {
	return errors.NewCriticalErrorf(
		"init Mail use case. %s is nil pointer.",
		name,
	).AddCode(http.StatusInternalServerError)
}

func newErrorConfig(name string) *errors.Error {
	return errors.NewCriticalErrorf(
		"init Mail use case. %s",
		name,
	).AddCode(http.StatusInternalServerError)
}
