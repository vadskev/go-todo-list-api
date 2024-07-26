package app

import (
	"log"

	"github.com/vadskev/go_final_project/internal/config"
	"github.com/vadskev/go_final_project/internal/config/env"
	"github.com/vadskev/go_final_project/internal/storage/db"
)

type configProvider struct {
	logConfig  config.LogConfig
	httpConfig config.HTTPConfig
	dbConfig   config.DBConfig
	repository db.Repository
	password   string
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

func (s *configProvider) DBConfig() config.DBConfig {
	if s.dbConfig == nil {
		cfg, err := env.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}
		s.dbConfig = cfg
	}
	return s.dbConfig
}

func (s *configProvider) DBRepository() db.Repository {
	if s.repository.DB() == nil {
		sqliteDb, err := db.NewRepository(s.DBConfig().Path())
		if err != nil {
			log.Fatalf("Failed to create db client: %v", err)
		}

		if err != nil {
			log.Fatalf("Ping error: %s", err.Error())
		}

		s.repository = sqliteDb
	}
	return s.repository
}

func (s *configProvider) GetPassFromEnv() string {
	if len(s.password) == 0 {
		cfg, err := env.NewPassConfig()
		if err != nil {
			log.Fatalf("failed to get password: %s", err.Error())
		}
		s.password = cfg.GetPass()
	}
	return s.password
}
