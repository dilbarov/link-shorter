package auth

import (
	authCommands "link-shorter/internal/auth/services/commands"
	userRepository "link-shorter/internal/user/repository"
)

type ServiceFacade struct {
	Commands *CommandBus
}

func NewAuthService(userRepo userRepository.Repository) *ServiceFacade {
	return &ServiceFacade{
		Commands: &CommandBus{
			Login: &authCommands.LoginCommandHandler{
				UserRepository: userRepo,
			},
			Register: &authCommands.RegisterCommandHandler{
				UserRepository: userRepo,
			},
		},
	}
}
