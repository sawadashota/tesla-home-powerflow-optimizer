package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
)

type (
	Middleware struct {
		r dependencies
	}
	dependencies interface {
		configuration.AppConfigProvider
	}
	Provider interface {
		Middleware() *Middleware
	}
)

func New(deps dependencies) *Middleware {
	return &Middleware{
		r: deps,
	}
}

func (m *Middleware) NewNOPMiddleware(next http.Handler) http.Handler {
	return next
}

func (m *Middleware) NewNoCacheMiddleware() func(http.Handler) http.Handler {
	return middleware.NoCache
}

func (m *Middleware) NewPanicRecoverMiddleware() func(http.Handler) http.Handler {
	return middleware.Recoverer
}
