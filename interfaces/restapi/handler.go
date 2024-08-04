package restapi

import (
	"net/http"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/app/usecase"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/restapi/errors"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/interfaces/restapi/middleware"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/pkg/restapi"
)

type (
	dependencies interface {
		configuration.AppConfigProvider
		usecase.Provider
	}
)

func NewHandler(r dependencies) http.Handler {
	m := middleware.New(r)
	return restapi.HandlerWithOptions(
		restapi.NewStrictHandlerWithOptions(
			r.Usecase(),
			[]restapi.StrictMiddlewareFunc{},
			restapi.StrictHTTPServerOptions{
				RequestErrorHandlerFunc:  errors.RequestErrorHandlerFunc,
				ResponseErrorHandlerFunc: errors.ResponseErrorHandlerFunc,
			},
		),
		restapi.ChiServerOptions{
			Middlewares: []restapi.MiddlewareFunc{
				m.NewLoggerMiddleware(),
				m.NewRequestValidatorMiddleware(),
			},
		},
	)
}
