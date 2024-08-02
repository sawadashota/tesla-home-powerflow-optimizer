package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate() error
}

func Validate(v any) error {
	if err := validator.New().Struct(v); err != nil {
		return fmt.Errorf("%w: %w", ErrValidationError, err)
	}
	return nil
}
