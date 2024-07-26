package env

import (
	"errors"
	"os"
)

const (
	PasswordEnvName = "TODO_PASSWORD"
)

type Password interface {
	GetPass() string
}

/**/

type passConfig struct {
	password string
}

func NewPassConfig() (Password, error) {
	pass := os.Getenv(PasswordEnvName)
	if len(pass) == 0 {
		return nil, errors.New("password not found")
	}

	return &passConfig{
		password: pass,
	}, nil
}

func (cfg *passConfig) GetPass() string {
	return cfg.password
}
