package link

import linkQueries "link-shorter/internal/link/services/queries"

type QueryBus struct {
	GetByHash *linkQueries.GetByHashQueryHandler
	GetById   *linkQueries.GetByIdQueryHandler
	GetAll    *linkQueries.GetAllQueryHandler
}
