package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	"link-shorter/internal/link/handlers"
	linkRepository "link-shorter/internal/link/repository"
	linkServices "link-shorter/internal/link/services"
	"link-shorter/pkg/db"
	"link-shorter/pkg/logger"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	log.Logger = logger.SetupLogger(conf.App.Env)

	database := db.NewDb(&conf.Db)

	router := http.NewServeMux()

	// Repositories
	linkRepo := linkRepository.NewPostgresLinkRepository(database)

	// Services
	linkService := linkServices.NewServiceFacade(linkRepo)

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

	log.Info().Msgf("Starting server on port %d", conf.App.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("")
	}
}
