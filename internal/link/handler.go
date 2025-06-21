package link

import (
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	"link-shorter/pkg/jwt"
	"link-shorter/pkg/middleware"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	LinkService *ServiceFacade
}

type Handler struct {
	*configs.Config
	LinkService *ServiceFacade
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		LinkService: deps.LinkService,
	}

	router.HandleFunc("GET /links", handler.getAll())
	router.HandleFunc("GET /links/{id}", handler.getById())
	router.HandleFunc("POST /links", handler.create())
	router.Handle("PATCH /links/{id}", middleware.IsAuthed(handler.update(), &deps.Config.Auth))
	router.HandleFunc("DELETE /links/{id}", handler.delete())
	router.HandleFunc("GET /r/{hash}", handler.goTo())
}

func (handler *Handler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hash := r.PathValue("hash")

		query := GetByHashQuery{
			Params: &GetByHashParams{
				Hash: hash,
			},
		}

		result, err := handler.LinkService.GetByHash(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		//err = handler.ClickRepository.Create(result.Id.String())
		//handler.EventBus.Publish()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, r, result.Url, http.StatusTemporaryRedirect)
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

		result, count, err := handler.LinkService.GetAll(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, NewPublicLinkList(result, &count), http.StatusOK)
	}
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
			Params: &GetByIDParams{
				Id: *id,
			},
		}

		result, err := handler.LinkService.GetById(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, NewPublicLink(result), http.StatusOK)
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

		result, err := handler.LinkService.Create(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicLink(result), http.StatusCreated)
	}
}

func (handler *Handler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user := ctx.Value(middleware.ContextUserKey).(*jwt.Data)

		log.Debug().Msgf("%v", user.Email)

		id, err := req.ParseUUID(r, "id")

		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		payload, err := req.HandleBody[UpdateRequest](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := UpdateCommand{
			Payload: &UpdatePayload{
				Id:   *id,
				Url:  payload.Url,
				Hash: payload.Hash,
			},
		}

		result, err := handler.LinkService.Update(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, NewPublicLink(result), http.StatusOK)
	}
}

func (handler *Handler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id, err := req.ParseUUID(r, "id")
		if err != nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		cmd := DeleteCommand{
			Payload: &GetByIDParams{
				Id: *id,
			},
		}

		_, err = handler.LinkService.Delete(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
