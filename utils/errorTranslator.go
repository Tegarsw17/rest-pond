package utils

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TranslateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)

	for _, fieldError := range validatorErrs {
		translation := fieldError.Translate(trans)
		errs = append(errs, translation)
	}
	return errs
}
