package auth

import authCommands "link-shorter/internal/auth/services/commands"

type CommandBus struct {
	Login    *authCommands.LoginCommandHandler
	Register *authCommands.RegisterCommandHandler
}
