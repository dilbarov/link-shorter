package user

import userCommands "link-shorter/internal/user/services/commands"

type CommandBus struct {
	Create *userCommands.CreateCommandHandler
	Update *userCommands.UpdateCommandHandler
	Delete *userCommands.DeleteCommandHandler
}
