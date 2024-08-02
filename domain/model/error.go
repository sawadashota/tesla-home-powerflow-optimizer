package model

import "errors"

var ErrValidationError = errors.New("validation error")

var (
	ErrGrantNotFound = errors.New("grant not found")
)
