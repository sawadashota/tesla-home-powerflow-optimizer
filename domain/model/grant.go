package model

import (
	"time"
)

type Grant struct {
	Subject      string    `validate:"required"`
	AccessToken  string    `validate:"required"`
	RefreshToken string    `validate:"required"`
	Scope        string    `validate:"required"`
	Expiry       time.Time `validate:"required"`
}

func (g *Grant) Validate() error {
	return Validate(g)
}
