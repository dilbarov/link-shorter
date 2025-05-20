package main

import (
	"errors"
	"fmt"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	"link-shorter/internal/link/handlers"
	"link-shorter/internal/link/repository"
	"link-shorter/internal/link/services"
	"link-shorter/pkg/db"
	"log"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	database := db.NewDb(&conf.Db)

	router := http.NewServeMux()

	// Repositories
	linkRepository := repository.NewPostgresLinkRepository(database)

	// Services
	linkService := link.NewServiceFacade(linkRepository)

	// Handler
	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})

	handlers.NewLinkHandler(router, handlers.HandlerDeps{
		Config:      conf,
		LinkService: linkService,
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port),
		Handler: router,
	}

	log.Printf("Starting server on port %d\n", conf.App.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ListenAndServe error: %v\n", err)
	}
}
