package app

import (
	"log"

	filmlibriary "github.com/Alphonnse/filmLibriary/internal/api/filmLibriary"
	"github.com/Alphonnse/filmLibriary/internal/api/user"
	"github.com/Alphonnse/filmLibriary/internal/config"
	"github.com/Alphonnse/filmLibriary/internal/repository"
	newone "github.com/Alphonnse/filmLibriary/internal/repository/sqlc"
	"github.com/Alphonnse/filmLibriary/internal/service"
	userService "github.com/Alphonnse/filmLibriary/internal/service/user"
	libriaryService "github.com/Alphonnse/filmLibriary/internal/service/filmLibriary"
)

type serviceProvider struct {
	serverConfig config.ServerConfig

	databaseConfig config.DatabaseConfig

	repository repository.Respository

	userService service.ServiceUserShape

	libriaryService service.ServiceLibriaryShape

	userImpl *user.ImplementationUser

	libriaryImpl *filmlibriary.ImplementationLibriary
}

func NewServiceProvider() *serviceProvider{
	return &serviceProvider{}
}

func (s *serviceProvider) ServerConfig() config.ServerConfig{
	if s.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			log.Fatalf("failed to get server config: %s", err.Error())
		}
		s.serverConfig = cfg
	}
	return s.serverConfig
}

func (s *serviceProvider) DatabaseConfig() config.DatabaseConfig{
	if s.databaseConfig == nil {
		cfg, err := config.NewDatabaseConfig()
		if err != nil {

		}
		s.databaseConfig = cfg
	}
	return s.databaseConfig
}

func (s *serviceProvider) Repository() repository.Respository{
	if s.repository == nil {
		s.repository = newone.NewRepository(
			s.DatabaseConfig(),
		)
	}
	return nil
}

func (s *serviceProvider) UserService() service.ServiceUserShape{
	if s.userService == nil {
		s.userService = userService.NewServiceUser(
			s.Repository(),
		)
	}
	return s.userService
}

func (s *serviceProvider) LibriaryService() service.ServiceLibriaryShape{
	if s.libriaryService == nil {
		s.libriaryService = libriaryService.NewServiceLibriary(
			s.Repository(),
		)
	}
	return s.libriaryService
}
