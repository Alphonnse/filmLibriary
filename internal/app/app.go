package app

import (
	"context"
	"log"
	"net/http"

	_ "github.com/Alphonnse/filmLibriary/docs"
	"github.com/Alphonnse/filmLibriary/internal/config"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Film Libriary RestAPI app
// @version 1.0
// @description Its the Film Libtiary app that uses pSQL, JWT netHttp lib

// @host localhost:8000
// @BasePath /

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
	a.serviceProvider.JwtConfig()
	a.serviceProvider.Repository()
	a.serviceProvider.UserService()
	a.serviceProvider.LibriaryService()
	a.serviceProvider.LibriaryImpl()
	a.serviceProvider.UserImpl()
	a.serviceProvider.AuthService()

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
	a.mux.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/swagger/doc.json")))
	a.mux.HandleFunc("/actor/addActorInfo", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.AddActorInfo))
	a.mux.HandleFunc("/actor/changeActorInfo", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.ChangeActorInfo))
	a.mux.HandleFunc("/actor/remove/", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.RmActorInfo))
	a.mux.HandleFunc("/actor/getList", a.serviceProvider.authService.JWTRegularUserAuth(a.serviceProvider.libriaryImpl.GetActorsInfoListWithFilms))
	a.mux.HandleFunc("/film/addFilmInfo", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.AddFilmInfo))
	a.mux.HandleFunc("/film/changeFilmInfo", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.ChangeFilmInfo))
	a.mux.HandleFunc("/film/remove/", a.serviceProvider.authService.JWTAdminAuth(a.serviceProvider.libriaryImpl.RmFilmInfo))
	a.mux.HandleFunc("/film/getList/", a.serviceProvider.authService.JWTRegularUserAuth(a.serviceProvider.libriaryImpl.GetFilmsList))
	a.mux.HandleFunc("/film/findFilm/", a.serviceProvider.authService.JWTRegularUserAuth(a.serviceProvider.libriaryImpl.GetFilmsListByFragment))
	a.mux.HandleFunc("/register", a.serviceProvider.userImpl.AddUser)
	a.mux.HandleFunc("/login", a.serviceProvider.userImpl.GetUser)
	a.mux.HandleFunc("/test/{some_shit}", a.serviceProvider.userImpl.Test)
}

type userInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
