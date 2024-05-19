package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	return &Validator{validator.New()}
}

func (v *Validator) Register(tag string, fn validator.Func) error {
	if err := v.RegisterValidation(tag, fn); err != nil {
		return validatorRegistrationTagError(tag, err)
	}

	return nil
}

func (v *Validator) ValidateStruct(data any) error {
	if err := v.Struct(data); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return internalError
		}
		for _, vErr := range err.(validator.ValidationErrors) {
			return validationErrorWithNameAndTag(vErr)
		}
	}
	return nil
}
