package validate

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetValidatorInstance() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}

// ValidateRequestBody validates the given struct or struct pointer
func ValidateRequestBody(s any) error {
	v := reflect.ValueOf(s)

	// If s is a pointer, get the value it points to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure s is a struct
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("ValidateRequestBody: argument must be a struct or a struct pointer")
	}

	// Get the validator instance and validate the struct
	return GetValidatorInstance().Struct(s)
}
