package validator

import (
	"exchange_rate/pkg/packages/errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	internalError = errors.New("Internal validation error")
)

func validatorRegistrationTagError(tag string, err error) *errors.Error {
	return errors.New(
		fmt.Sprintf(
			"Validator registration error at tag {%s}. Error: %v",
			tag,
			err),
	)
}

func validationErrorWithNameAndTag(err validator.FieldError) *errors.Error {
	return errors.New(
		fmt.Sprintf(
			"Validation error on field:%s with tag:%s", err.Field(), err.Tag(),
		),
	)
}
