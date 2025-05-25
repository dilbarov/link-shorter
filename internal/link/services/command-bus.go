package link

import (
	linkCommands "link-shorter/internal/link/services/commands"
)

type CommandBus struct {
	Create *linkCommands.CreateCommandHandler
	Update *linkCommands.UpdateCommandHandler
	Delete *linkCommands.DeleteCommandHandler
}
