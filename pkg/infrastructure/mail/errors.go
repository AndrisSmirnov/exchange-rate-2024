package mail

import "exchange_rate/pkg/packages/errors"

var ErrMailConfig = errors.New("error in mail. empty config")

func errSendMail(err error) *errors.Error {
	return errors.Newf("error sending mail. Reason: %v", err)
}

func newErrorNoEnvVar(variable string) *errors.Error {
	return errors.NewCriticalErrorf(
		"error no environment variable %s.",
		variable,
	)
}
