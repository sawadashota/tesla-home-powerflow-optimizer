package driver

import (
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type (
	ServerRegistry interface {
		logx.Provider
	}
	serverRegistry struct {
		EssentialRegistry
	}
)

var _ ServerRegistry = new(serverRegistry)

func NewServerRegistry() (ServerRegistry, error) {
	registry, err := NewEssentialRegistry()
	if err != nil {
		return nil, err
	}
	return &serverRegistry{
		EssentialRegistry: registry,
	}, nil
}
