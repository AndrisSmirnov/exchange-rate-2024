package errors

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	SECTION_SEPARATOR = " :: "
	// TODO: DEPRECATED. Delete after refactoring.
	TEXT_DATA_ERROR = "data error"
	// TODO: DEPRECATED. Delete after refactoring.
	TEXT_INTERNAL_ERROR = "internal error"
	// TODO: DEPRECATED. Delete after refactoring.
	TEXT_CRITICAL_ERROR = "critical error"
)

type ErrorType byte

const (
	DataError ErrorType = iota + 1
	InternalError
	CriticalError
)

type Error struct {
	Type          ErrorType
	Message       string
	ExtendedError error
	Stack         []byte
	Code          int
}

type ErrorTransport struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

func (err *Error) Error() string {
	if err == nil {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			logrus.Warnf("called from %s\n", details.Name())
		}

		return fmt.Sprintf(
			"%s%s%s",
			TEXT_CRITICAL_ERROR,
			SECTION_SEPARATOR,
			"error is nil",
		)
	}

	t := err.typeAndCodeToString()
	if err.ExtendedError == nil {
		return fmt.Sprintf(
			"%s%s%s",
			t,
			SECTION_SEPARATOR,
			err.Message,
		)
	} else {
		return fmt.Sprintf(
			"%s%s%s%s%s",
			t,
			SECTION_SEPARATOR,
			err.Message,
			SECTION_SEPARATOR,
			err.ExtendedError.Error(),
		)
	}
}

func (t ErrorType) ToCode() int {
	switch t {
	case DataError:
		return http.StatusBadRequest
	case InternalError:
		return http.StatusInternalServerError
	case CriticalError:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}
