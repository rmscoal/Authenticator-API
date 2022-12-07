package v1

import (
	"regexp"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func phoneNumValidator(fl validator.FieldLevel) bool {
	// Regex according to Indonesia phone numbers format.
	regex, _ := regexp.Compile(`^(\+62|62|021|0274)?[\s-]?0?8?[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)

	return regex.MatchString(fl.Field().String())
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// Add a custom phone regex validation into Validator.
	cv.Validator.RegisterValidation("phone", phoneNumValidator)

	if err := cv.Validator.Struct(i); err != nil {
		// Add a custom error message regarding its field.
		return &echo.HTTPError{
			Message: "invalid " + err.(validator.ValidationErrors)[0].Field(),
		}
	}
	return nil
}
