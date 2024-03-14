package app

import (
	"context"
	"log"
	"net/http"

	"github.com/Alphonnse/filmLibriary/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	server          *http.Server
	mux             *http.ServeMux
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	log.Printf("Running server on: %s", a.serviceProvider.serverConfig.Address())
	return a.runServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	log.Println("dependencies are satisfied")
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
    a.serviceProvider = NewServiceProvider()

    a.serviceProvider.ServerConfig()
    a.serviceProvider.DatabaseConfig()
    a.serviceProvider.Repository()
    a.serviceProvider.UserService()
    a.serviceProvider.LibriaryService()
	a.serviceProvider.LibriaryImpl()
	a.serviceProvider.UserImpl()

	log.Println("Service provider initialized")
    return nil
}

func (a *App) initServer(_ context.Context) error {
	a.mux = http.NewServeMux()
	a.server = &http.Server{
		Addr:    a.serviceProvider.ServerConfig().Address(),
		Handler: a.mux,
	}
	return nil
}

func (a *App) runServer() error {
	a.setupRoutes()
	err := a.server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error server listening")
	}
	return err 
}

func (a *App) setupRoutes() {
	a.mux.HandleFunc("/", a.serviceProvider.libriaryImpl.AddActorInfo)
}
