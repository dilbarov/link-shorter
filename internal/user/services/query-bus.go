package user

import userQueries "link-shorter/internal/user/services/queries"

type QueryBus struct {
	GetById    *userQueries.GetByIdQueryHandler
	GetByEmail *userQueries.GetByEmailQueryHandler
	GetAll     *userQueries.GetAllQueryHandler
}
