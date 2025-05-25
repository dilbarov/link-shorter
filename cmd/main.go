package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	linkHandlers "link-shorter/internal/link/handlers"
	linkRepository "link-shorter/internal/link/repository"
	linkServices "link-shorter/internal/link/services"
	userHandlers "link-shorter/internal/user/handlers"
	userRepository "link-shorter/internal/user/repository"
	userServices "link-shorter/internal/user/services"
	"link-shorter/pkg/db"
	"link-shorter/pkg/logger"
	"link-shorter/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	log.Logger = logger.SetupLogger(conf.App.Env)

	database := db.NewDb(&conf.Db)

	router := http.NewServeMux()

	// Repositories
	linkRepo := linkRepository.NewPostgresLinkRepository(database)
	userRepo := userRepository.NewPostgresUserRepository(database)

	// Services
	linkService := linkServices.NewServiceFacade(linkRepo)
	userService := userServices.NewServiceFacade(userRepo)

	// Handler
	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})

	linkHandlers.NewLinkHandler(router, linkHandlers.LinkHandlerDeps{
		Config:      conf,
		LinkService: linkService,
	})

	userHandlers.NewUserHandler(router, userHandlers.UserHandlerDeps{
		Config:      conf,
		UserService: userService,
	})

	// Middlewares
	loggingMiddleware := middleware.NewLoggingMiddleware(log.Logger)
	stack := middleware.Chain(middleware.CORS, loggingMiddleware)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port),
		Handler: stack(router),
	}

	log.Info().Msgf("Starting server on port %d", conf.App.Port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("")
	}
}
