package driver

import (
	"github.com/sawadashota/tesla-home-powerflow-optimizer/domain/repository"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/driver/configuration"
)

type (
	configurationProvider configuration.Provider
	repositoryProvider    repository.Provider
)
