package config

import (
	"log"
	"os"
)

const (
	secretEnv = "SECRET"
)

type JwtConfig interface {
	Secret() []byte
}

type jwtConfig struct {
	secret []byte	
}

func NewJwtConfig() (*jwtConfig, error) {
	secret := os.Getenv(secretEnv)
	if secret == "" {
		log.Fatal("secret is not found in the environment")
	}
	return &jwtConfig{
		secret: []byte(secret),
	}, nil
}

func (cfg *jwtConfig) Secret() []byte {
	return cfg.secret
}
