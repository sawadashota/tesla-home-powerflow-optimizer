package tesla

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/httpclient"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

const userAgent = "TeslaHomePowerflowOptimizer/1.0"

type (
	Client struct {
		r           dependencies
		tokenSource oauth2.TokenSource
		httpclient  *http.Client
	}
	dependencies interface {
		logx.Provider
		configuration.AppConfigProvider
		configuration.TeslaOAuthConfigProvider
		httpclient.Provider
		repository.GrantRepositoryProvider
	}
)

func NewClient(ctx context.Context, r dependencies) repository.VehicleRepository {
	source := newTokenSource(ctx, r)
	return &Client{
		r:           r,
		tokenSource: source,
		httpclient:  oauth2.NewClient(ctx, source),
	}
}
