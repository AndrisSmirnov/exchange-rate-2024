package errors

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsStringToError(errString string) *Error {
	arg := strings.Split(errString, SECTION_SEPARATOR)
	if len(arg) < 2 {
		return NewInternalError(errString)
	}

	errT, code := typeAndCodeToInternal(arg[0])
	err := newError(errT, arg[1]).AddCode(code)
	if len(arg) > 2 {
		subErrStart := len(arg[0]) + len(arg[1]) + (len(SECTION_SEPARATOR) * 2)
		err.ExtendedError = ParsStringToError(
			errString[subErrStart:],
		)
	}

	return err
}

func (err *Error) ParseCriticalError() string {
	return fmt.Sprintf(
		"Type: %v | Error: %v | Reason: %v",
		err.Type,
		err.Message,
		err.ExtendedError,
	)
}

func (e *Error) typeAndCodeToString() string {
	return fmt.Sprintf("%02d%03d", e.Type, e.Code)
}

func typeAndCodeToInternal(str string) (ErrorType, int) {
	if len(str) != 5 {
		return typeAndCodeToInternalDep(str)
	}

	var errT ErrorType
	t, err := strconv.Atoi(str[:2])
	if err == nil {
		errT = ErrorType(t)
	} else {
		errT = InternalError
	}

	var errC int
	code, err := strconv.Atoi(str[2:])
	if err == nil {
		errC = code
	} else {
		errC = 0
	}

	return errT, errC
}

// typeAndCodeToInternalDep - convert type string to type int
// Is used for backward compatible with package versions <= v0.1.3
// TODO: Delete after refactoring.
func typeAndCodeToInternalDep(str string) (ErrorType, int) {
	switch str {
	case TEXT_DATA_ERROR:
		return DataError, 0
	case TEXT_INTERNAL_ERROR:
		return InternalError, 0
	case TEXT_CRITICAL_ERROR:
		return CriticalError, 0
	default:
		return InternalError, 0
	}
}
