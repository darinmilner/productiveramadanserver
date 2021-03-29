package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/formurl", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("Got an invalid form when should be valid")
	}
}

func TestFormRequired(t *testing.T) {
	r := httptest.NewRequest("POST", "/formurl", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Form shows valid when required fields are missing")
	}

}

func TestHasARequiredField(t *testing.T) {
	r := httptest.NewRequest("Post", "/whatever", nil)

	form := New(r.PostForm)

	field := "whatever"

	has := form.HasARequiredField(field)

	if has {
		t.Error("Form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	has = form.HasARequiredField("a")

	if !has {
		t.Error("Shows form does not have existing field")
	}

}

func TestFormHasMinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	field := "testField"

	form.MinLength("a", 3)
	if form.Valid() {
		t.Error("Form does not return minlength")
	}
	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("Should have a form error but didn't get one")
	}

	postedData := url.Values{}

	postedData.Add(field, "a4213213213")

	form = New(postedData)

	form.MinLength(field, 30)
	if form.Valid() {
		t.Error("Form does not check minlength")
	}

	postedData = url.Values{}
	postedData.Add("another-field", "abc123")
	form = New(postedData)
	form.MinLength("another-field", 2)

	if !form.Valid() {
		t.Error("Form field is larger than min value")
	}

	isError = form.Errors.Get("another-field")
	if isError != "" {
		t.Error("Should not have an error but got one")
	}

}

func TestIsEmail(t *testing.T) {
	invalidEmail := "abc"
	postedValues := url.Values{}

	form := New(postedValues)

	form.IsEmail(invalidEmail)

	if form.Valid() {
		t.Error("Form shows valid email for non-existent field")
	}

	field := "email"
	validEmail := "abc@abc.com"
	postedValues = url.Values{}
	postedValues.Add(field, validEmail)

	form = New(postedValues)

	form.IsEmail(field)

	if !form.Valid() {
		t.Error("Got an invalid email when should be valid")
	}

	postedValues = url.Values{}
	postedValues.Add(field, invalidEmail)

	form = New(postedValues)
	form.IsEmail(field)
	if form.Valid() {
		t.Error("Got a valid email when should be invalid")
	}

}
