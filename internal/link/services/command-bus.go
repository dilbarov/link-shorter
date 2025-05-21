package link

import (
	commands "link-shorter/internal/link/services/commands"
)

type CommandBus struct {
	Create *commands.CreateCommandHandler
	Update *commands.UpdateCommandHandler
}
