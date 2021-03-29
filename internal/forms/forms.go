package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

//Valid returns true if the form field has no errors
func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Required checks for required fields and sends a message if empty
func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field can not be empty")
		}
	}
}

//HasARequiredField checks for required fields
func (form *Form) HasARequiredField(field string) bool {
	formText := form.Get(field)

	if formText == "" {
		return false
	}
	return true
}

//MinLength check if the field text is a min length
func (form *Form) MinLength(field string, length int) bool {
	formText := form.Get(field)

	if len(formText) < length {
		form.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters", length))
		return false
	}

	return true
}

func (form *Form) IsEmail(field string) {
	if !govalidator.IsEmail(form.Get(field)) {
		form.Errors.Add(field, "Please enter a valid email")
	}
}
