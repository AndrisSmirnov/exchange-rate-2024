package scheduler_service

import (
	"exchange_rate/pkg/packages/errors"
)

func newErrorNoEnvVar(variable string) *errors.Error {
	return errors.NewCriticalErrorf(
		"error no environment variable %s.",
		variable,
	)
}
