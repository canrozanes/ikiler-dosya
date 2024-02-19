package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ikiler-dosya/pkg/router"

	"github.com/unrolled/secure"
)

type Config struct {
	ServerOrigin  string
	SecureOptions secure.Options
	Audience      string
	Auth0Domain   string
}

type App struct {
	Config Config
}

func (app *App) RunServer() {
	router := router.Router(app.Config.ServerOrigin, app.Config.Audience, app.Config.Auth0Domain)

	router = secure.New(app.Config.SecureOptions).Handler(router)

	server := &http.Server{
		Addr:    app.Config.ServerOrigin,
		Handler: router,
	}

	log.Printf("API server listening on %s", server.Addr)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("API server closed: err: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("got shutdown signal. shutting down server...")

	localCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(localCtx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("server shutdown complete")
}
