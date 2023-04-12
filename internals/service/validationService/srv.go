package validationService

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type validationSrv struct {
}

type ValidationService interface {
	Validate(data any) error
}

func NewValidationSrv() ValidationService {
	return &validationSrv{}
}

func (t *validationSrv) Validate(data any) error {

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				return fmt.Errorf("%s is required", err.Field())
			case "email":
				return fmt.Errorf("%s is not a valid email address", err.Field())
			default:
				return fmt.Errorf("%s is invalid", err.Field())
			}
		}
	}

	return validate.Struct(data)
}
