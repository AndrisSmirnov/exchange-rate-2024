package vo

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UUID string

func (u UUID) String() string {
	return string(u)
}

func NewID() UUID {
	return UUID(uuid.NewString())
}

func ValidateUUID(fl validator.FieldLevel) bool {
	validationValue, ok := fl.Field().Interface().(UUID)
	if !ok {
		logrus.Warn("error to convert vo.UUID in validator", fl.Field().Interface())
		return false
	}
	return validationValue.Validate()
}

func (u UUID) Validate() bool {
	_, err := uuid.Parse(u.String())
	if err != nil {
		return false
	}

	return true
}
