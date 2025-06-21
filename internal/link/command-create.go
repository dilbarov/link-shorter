package link

import "context"

const CreateLinkCommandName = "link.create"

type CreateCommand struct {
	Payload *CreatePayload
}

func (c CreateCommand) Name() string {
	return CreateLinkCommandName
}

type CreateCommandHandler struct {
	LinkRepository Repository
}

func NewCreateCommandHandler(repo Repository) *CreateCommandHandler {
	return &CreateCommandHandler{LinkRepository: repo}
}

func (h *CreateCommandHandler) Handle(ctx context.Context, cmd CreateCommand) (*Model, error) {
	link := NewLink(cmd.Payload.Url)
	for {
		existsLink, _ := h.LinkRepository.GetByHash(link.Hash)
		if existsLink == nil {
			break
		}
		link.GenerateHash()
	}
	createdLink, err := h.LinkRepository.Create(link)
	if err != nil {
		return nil, err
	}
	return createdLink, nil
}
