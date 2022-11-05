package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

const webPort = 80

type App struct{}

func main() {
	app := &App{}

	log.Printf("Starting broker service on port %d\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(errors.Wrap(err, "failed to start broker service"))
	}
}
