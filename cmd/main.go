package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	"link-shorter/internal/click"
	"link-shorter/internal/link"
	"link-shorter/internal/user"
	"link-shorter/pkg/cqrs"
	"link-shorter/pkg/db"
	"link-shorter/pkg/jwt"
	"link-shorter/pkg/logger"
	"link-shorter/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	log.Logger = logger.SetupLogger(conf.App.Env)
	database := db.NewDb(&conf.Db)
	router := http.NewServeMux()
	eventBus := cqrs.NewInMemoryEventBus()

	// Repositories
	linkRepo := link.NewPostgresLinkRepository(database)
	userRepo := user.NewPostgresUserRepository(database)
	clickRepo := click.NewPostgresClickRepository(database)

	// Services
	userService := user.NewServiceFacade(userRepo)
	jwtService := jwt.NewJWTService(conf.Auth.Secret)
	authService := auth.NewAuthService(userRepo, jwtService)
	linkService := link.NewServiceFacade(link.ServiceFacadeDeps{
		LinkRepository: linkRepo,
	})
	_ = click.NewServiceFacade(click.ServiceFacadeDeps{
		EventBus:        eventBus,
		ClickRepository: clickRepo,
	})

	// Handlers
	auth.NewHandler(router, auth.HandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.HandlerDeps{
		Config:      conf,
		LinkService: linkService,
	})

	user.NewUserHandler(router, user.HandlerDeps{
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
