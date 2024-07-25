package app

import (
	"log"

	"github.com/vadskev/go_final_project/internal/config"
	"github.com/vadskev/go_final_project/internal/config/env"
)

type configProvider struct {
	logConfig  config.LogConfig
	httpConfig config.HTTPConfig
}

func newConfigProvider() *configProvider {
	return &configProvider{}
}

func (s *configProvider) LogConfig() config.LogConfig {
	if s.logConfig == nil {
		cfg, err := env.NewLogConfig()
		if err != nil {
			log.Fatalf("failed to get log config: %s", err.Error())
		}
		s.logConfig = cfg
	}
	return s.logConfig
}

func (s *configProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}
