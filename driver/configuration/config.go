package configuration

type (
	Provider interface {
		Configuration() Configuration
	}
	Configuration interface {
		AppConfigProvider
		ServerConfigProvider
		TeslaOAuthConfigProvider
		TeslaAPIConfigProvider
	}
	configuration struct {
		AppConfigProvider
		ServerConfigProvider
		TeslaOAuthConfigProvider
		TeslaAPIConfigProvider
	}
)

func New() (Configuration, error) {
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

	teslaAPIConfig, err := NewTeslaAPIConfig()
	if err != nil {
		return nil, err
	}

	return &configuration{
		AppConfigProvider:        NewAppConfigProvider(appConfig),
		ServerConfigProvider:     NewServerConfigProvider(serverConfig),
		TeslaOAuthConfigProvider: NewTeslaOAuthConfigProvider(teslaOAuthConfig),
		TeslaAPIConfigProvider:   NewTeslaAPIConfigProvider(teslaAPIConfig),
	}, nil
}
