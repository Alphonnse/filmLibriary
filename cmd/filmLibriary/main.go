package main

import (
	"context"
	"log"

	"github.com/Alphonnse/filmLibriary/internal/app"
)


// @title Film Libriary RestAPI app
// @version 1.0
// @description Its the Film Libtiary app that uses pSQL, JWT netHttp lib

// @host localhost:8080
// @BasePath /


func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

}
