package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"

	"github.com/nullifidianz/go-crud/src/config/catch_err"
)

var (
	Validate  = validator.New()
	translate ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		translate, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, translate)
	}
}

func ValidateUserError(
	validation_error error) *catch_err.MyError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return catch_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorsCause := []catch_err.Cause{}
		for _, err := range jsonValidationError {
			cause := catch_err.Cause{
				Field:   err.Field(),
				Message: err.Translate(translate),
			}
			errorsCause = append(errorsCause, cause)
		}

		return catch_err.NewBadRequestError("some of the fields are missing or invalid, please check and try again")

	} else if validation_error != nil {
		return catch_err.NewBadRequestError("Error: " + validation_error.Error())

	} else {
		return catch_err.NewBadRequestError("unknown error")
	}
}
