package link

import "link-shorter/internal/link/services/queries"

type QueryBus struct {
	GetByHash *queries.GetByHashQueryHandler
	GetById   *queries.GetByIdQueryHandler
}
