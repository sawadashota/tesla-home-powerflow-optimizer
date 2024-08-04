package model

import (
	"github.com/morikuni/failure/v2"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate() error
}

func Validate(v any) error {
	if err := validator.New().Struct(v); err != nil {
		return failure.New(
			ErrCodeValidationError,
			failure.Message("validation error"),
			failure.CallStackOf(err),
		)
	}
	return nil
}
