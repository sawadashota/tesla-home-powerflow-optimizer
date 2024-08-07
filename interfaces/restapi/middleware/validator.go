package middleware

import (
	"net/http"

	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/pkg/restapi"
)

func (m *Middleware) NewRequestValidatorMiddleware() func(http.Handler) http.Handler {
	swagger, err := restapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	return nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, &nethttpmiddleware.Options{
		SilenceServersWarning: true,
	})
}
