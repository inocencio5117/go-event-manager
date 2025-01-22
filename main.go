package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/inocencio5117/go-event-manager/config"
	_ "github.com/inocencio5117/go-event-manager/docs"
	"github.com/inocencio5117/go-event-manager/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Go Event Manager API Specification
// @version		0.1
// @description	This is the API specification for the Go Event Manager.
func main() {
	if err := run(); err != nil {
		log.Fatal("Error starting server", err)
		os.Exit(1)
	}
}

func run() error {
	config.ConnectDatabase()
	config.ConnectRedis()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	routes.EventRoutes(r)
	routes.AttendeeRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	log.Println("Server running on :8080")

	return nil
}
