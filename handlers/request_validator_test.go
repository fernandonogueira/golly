package handlers

import (
	"github.com/fernandonogueira/golly/models"
	"testing"
	"reflect"
)

func TestNewRequestValidator(t *testing.T) {
	validator := NewRequestValidator();
	t.Log(reflect.TypeOf(validator))

	if validator == nil {
		t.Error("Validator should not be nil")
	}
}

func TestValidator(t *testing.T) {
	validator := NewRequestValidator()

	request := models.GollyRequest{}
	validationErrors := validator.Validate(request, true)

	if validationErrors.Error == "" {
		t.Error("Error should not be empty because " +
			"request do not have webhook " +
			"address and is async = true. " +
			"Got: ", validationErrors)
	}
}