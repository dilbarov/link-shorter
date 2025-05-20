package link

import (
	commands "link-shorter/internal/link/services/commands"
)

type CommandBus struct {
	CreateHandler *commands.CreateHandler
	UpdateHandler *commands.UpdateHandler
}
