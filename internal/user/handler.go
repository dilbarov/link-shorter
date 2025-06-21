package user

import (
	"link-shorter/configs"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	UserService *ServiceFacade
}

type Handler struct {
	*configs.Config
	UserService *ServiceFacade
}

func NewUserHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		UserService: deps.UserService,
	}

	router.HandleFunc("GET /users", handler.getAll())
	router.HandleFunc("GET /users/{id}", handler.getById())
	router.HandleFunc("POST /users", handler.create())
	router.HandleFunc("PATCH /users/{id}", handler.update())
	router.HandleFunc("DELETE /users/{id}", handler.delete())
}

func (handler *Handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, err := req.ParseUUID(r, "id")

		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		query := GetByIdQuery{
			Params: &GetByIdParams{
				Id: *id,
			},
		}

		result, err := handler.UserService.GetById(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicUser(result), http.StatusOK)
	}
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		filter := GetAllParams{}
		err := req.ParseQuery(r, &filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := GetAllQuery{
			Params: &filter,
		}

		users, count, err := handler.UserService.GetAll(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicUserList(users, &count), http.StatusOK)
	}
}

func (handler *Handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := req.HandleBody[CreatePayload](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := CreateCommand{
			Payload: payload,
		}

		result, err := handler.UserService.Create(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicUser(result), http.StatusCreated)
	}
}

func (handler *Handler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, err := req.ParseUUID(r, "id")
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		payload, err := req.HandleBody[UpdatePayload](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := UpdateCommand{
			Payload: &UpdatePayload{
				Id:   *id,
				Name: payload.Name,
			},
		}

		result, err := handler.UserService.Update(ctx, cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicUser(result), http.StatusAccepted)
	}
}

func (handler *Handler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ParseUUID(r, "id")
		ctx := r.Context()
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		cmd := DeleteCommand{
			Payload: &GetByIdParams{
				Id: *id,
			},
		}

		err = handler.UserService.Delete(ctx, cmd)
	}
}
