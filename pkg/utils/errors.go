package utils

import (
	"exchange_rate/pkg/packages/errors"
	"fmt"
)

var (
	errorUnexpectedType = errors.New("unexpected type for envVal")
	errorParseEnv       = errors.New(
		"unable to parse type of environment",
	)
)

func newErrParse(variable, varType, reason string) *errors.Error {
	return errors.New(
		fmt.Sprintf("environment variable not found. Variable: %s. Error on parse %s. Reason: %s",
			variable,
			varType,
			reason,
		),
	)
}

func newErrNoEnv(variable string) *errors.Error {
	return errors.New(
		fmt.Sprintf("environment variable not found. Variable: %s",
			variable,
		),
	)
}
