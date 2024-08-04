package configuration

type (
	Provider interface {
		AppConfigProvider
		ServerConfigProvider
		TeslaOAuthConfigProvider
	}
	//Configuration interface {
	//	AppConfigProvider
	//	ServerConfigProvider
	//	TeslaOAuthConfigProvider
	//}
	configuration struct {
		AppConfigProvider
		ServerConfigProvider
		TeslaOAuthConfigProvider
	}
)

func New() (Provider, error) {
	appConfig, err := NewAppConfig()
	if err != nil {
		return nil, err
	}

	serverConfig, err := NewServerConfig()
	if err != nil {
		return nil, err
	}

	teslaOAuthConfig, err := NewTeslaOAuthConfig()
	if err != nil {
		return nil, err
	}

	return &configuration{
		AppConfigProvider:        NewAppConfigProvider(appConfig),
		ServerConfigProvider:     NewServerConfigProvider(serverConfig),
		TeslaOAuthConfigProvider: NewTeslaOAuthConfigProvider(teslaOAuthConfig),
	}, nil
}
