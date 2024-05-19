package errors

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func ErrLoggerAppWarning(err error) {
	b, err := json.Marshal(err)
	if err == nil {
		logrus.Warning("[!] App internal error: ", string(b))
	} else {
		logrus.Warning("[!] App internal error: ", err)
	}
}

func ErrLoggerAppError(err error) {
	b, err := json.Marshal(err)
	if err == nil {
		logrus.Error("[!] App internal error: ", string(b))
	} else {
		logrus.Error("[!] App internal error: ", err)
	}
}
