package mail

import "exchange_rate/pkg/packages/errors"

var (
	ErrMailConfig            = errors.New("error in mail. empty config")
	ErrCantUploadTemplate    = errors.New("error in upload the template file")
	ErrCantSetDataToTemplate = errors.New("error in set data into template file")
)

func errSendMail(err error) *errors.Error {
	return errors.Newf("error sending mail. Reason: %v", err)
}

func newErrorNoEnvVar(variable string) *errors.Error {
	return errors.NewCriticalErrorf(
		"error no environment variable %s.",
		variable,
	)
}
