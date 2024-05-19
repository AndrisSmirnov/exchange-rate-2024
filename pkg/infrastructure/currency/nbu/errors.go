package nbu

import "exchange_rate/pkg/packages/errors"

var (
	errConfigNotValid = errors.New("config not valid")
)

func errToCreateConfig(err error) *errors.Error {
	return errors.New("failed to create NBU currency config with base link. Reason: " + err.Error())
}

func newErrorNoEnvVar(variable string) *errors.Error {
	return errors.NewCriticalErrorf(
		"error no environment variable %s.",
		variable,
	)
}
