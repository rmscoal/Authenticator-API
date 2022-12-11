package v1

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

const (
	_defaultMininumIntCount       = 2
	_deafultUpperCaseMinimumCount = 1
)

type CustomValidator struct {
	Validator *validator.Validate
}

// phoneNumValidator verifies whether the given phone number
// satisfies the Indonesian phone number format that is done
// using a regex.
func phoneNumValidator(fl validator.FieldLevel) bool {
	regex, _ := regexp.Compile(`^(\+62|62|021|0274)?[\s-]?0?8?[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)

	return regex.MatchString(fl.Field().String())
}

// passwordValidator verifies whether the given password has
// satisfies the minimum amount of integers and uppercase
// letter contained in the password.
func passwordValidator(fl validator.FieldLevel) bool {
	var intCount, upperCaseCount int = 0, 0
	var psw string = fl.Field().String()

	for i := 0; i < len(psw); i++ {
		if unicode.IsDigit(rune(psw[i])) {
			intCount++
		}

		if unicode.IsLetter(rune(psw[i])) {
			if string(psw[i]) == strings.ToUpper(string(psw[i])) {
				upperCaseCount++
			}
		}
	}

	return intCount >= _defaultMininumIntCount && upperCaseCount >= _deafultUpperCaseMinimumCount
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// Register custom validators here:
	cv.Validator.RegisterValidation("phone", phoneNumValidator)
	cv.Validator.RegisterValidation("password", passwordValidator)

	if err := cv.Validator.Struct(i); err != nil {
		// Add a custom error message regarding its field.
		return &echo.HTTPError{
			Message: "invalid " + err.(validator.ValidationErrors)[0].Field(),
		}
	}
	return nil
}
