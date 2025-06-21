package auth

import (
	"context"
	"link-shorter/internal/user"
	"link-shorter/pkg/jwt"
)

type ServiceFacade struct {
	Login    func(ctx context.Context, cmd LoginCommand) (string, error)
	Register func(ctx context.Context, cmd RegisterCommand) (string, error)
}

func NewAuthService(userRepo user.Repository, jwtService *jwt.Service) *ServiceFacade {
	loginHandler := NewLoginCommandHandler(userRepo, jwtService)
	registerHandler := NewRegisterCommandHandler(jwtService, userRepo)

	return &ServiceFacade{
		Login:    loginHandler.Handle,
		Register: registerHandler.Handle,
	}
}
