package auth

import (
	authCommands "link-shorter/internal/auth/services/commands"
	userRepository "link-shorter/internal/user/repository"
	"link-shorter/pkg/jwt"
)

type ServiceFacade struct {
	Commands *CommandBus
}

func NewAuthService(userRepo userRepository.Repository, jwtService *jwt.Service) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			Login: &authCommands.LoginCommandHandler{
				UserRepository: userRepo,
				JwtService:     jwtService,
			},
			Register: &authCommands.RegisterCommandHandler{
				UserRepository: userRepo,
				JwtService:     jwtService,
			},
		},
	}
}
