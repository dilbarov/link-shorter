package handlers

import (
	"link-shorter/configs"
	userPayloads "link-shorter/internal/user/payloads"
	userResponses "link-shorter/internal/user/responses"
	userServices "link-shorter/internal/user/services"
	userCommands "link-shorter/internal/user/services/commands"
	userQueries "link-shorter/internal/user/services/queries"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type UserHandlerDeps struct {
	*configs.Config
	UserService *userServices.ServiceFacade
}

type UserHandler struct {
	*configs.Config
	UserService *userServices.ServiceFacade
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		Config:      deps.Config,
		UserService: deps.UserService,
	}

	router.HandleFunc("GET /users", handler.getAll())
	router.HandleFunc("GET /users/{id}", handler.getById())
	router.HandleFunc("POST /users", handler.create())
	router.HandleFunc("PATCH /users/{id}", handler.update())
	router.HandleFunc("DELETE /users/{id}", handler.delete())
}

func (handler *UserHandler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ParseUUID(r, "id")

		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		query := userQueries.GetByIdQuery{
			Params: &userPayloads.GetByIdParams{
				Id: *id,
			},
		}

		result, err := handler.UserService.Queries.GetById.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, userResponses.NewPublicUser(result), http.StatusOK)
	}
}

func (handler *UserHandler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := userPayloads.GetAllParams{}
		err := req.ParseQuery(r, &filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := userQueries.GetAllQuery{
			Params: &filter,
		}

		users, count, err := handler.UserService.Queries.GetAll.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, userResponses.NewPublicUserList(users, &count), http.StatusOK)
	}
}

func (handler *UserHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[userPayloads.CreatePayload](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := userCommands.CreateCommand{
			Payload: payload,
		}

		result, err := handler.UserService.Commands.Create.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, userResponses.NewPublicUser(result), http.StatusCreated)
	}
}

func (handler *UserHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ParseUUID(r, "id")
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		payload, err := req.HandleBody[userPayloads.UpdatePayload](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := userCommands.UpdateCommand{
			Payload: &userPayloads.UpdatePayload{
				Id:   *id,
				Name: payload.Name,
			},
		}

		result, err := handler.UserService.Commands.Update.Execute(cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, userResponses.NewPublicUser(result), http.StatusAccepted)
	}
}

func (handler *UserHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ParseUUID(r, "id")
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		cmd := userCommands.DeleteCommand{
			Payload: &userPayloads.GetByIdParams{
				Id: *id,
			},
		}

		err = handler.UserService.Commands.Delete.Execute(cmd)
	}
}
