package config

import (
	"errors"
	"net"
	"os"
)

const (
	serverHostEnvName = "SERVER_HOST" 
	serverPortEnvName = "SERVER_PORT" 
)

type ServerConfig interface {
	Address() string
}

type serverConfig struct {
	host string
	port string
}

func NewServerConfig() (ServerConfig, error) {
	host := os.Getenv(serverHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("Server host not found")
	}

	port := os.Getenv(serverPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("Server port not found")
	}

	return &serverConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
