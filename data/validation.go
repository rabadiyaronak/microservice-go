package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSku)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {

	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrors ValidationErrors
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}

		returnErrors = append(returnErrors, ve)
	}

	return returnErrors
}

func validateSku(fl validator.FieldLevel) bool {

	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	match := reg.FindAllString(fl.Field().String(), -1)

	return len(match) == 1
}
