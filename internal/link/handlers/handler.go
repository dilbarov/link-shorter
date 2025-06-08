package link

import (
	"github.com/rs/zerolog/log"
	"link-shorter/configs"
	linkPayloads "link-shorter/internal/link/payloads"
	linkResponses "link-shorter/internal/link/responses"
	linkServices "link-shorter/internal/link/services"
	linkCommands "link-shorter/internal/link/services/commands"
	linkQueries "link-shorter/internal/link/services/queries"
	cqrsInterfaces "link-shorter/pkg/cqrs/interfaces"
	"link-shorter/pkg/jwt"
	"link-shorter/pkg/middleware"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	LinkService *linkServices.ServiceFacade
	EventBus    cqrsInterfaces.EventBus
}

type Handler struct {
	*configs.Config
	LinkService *linkServices.ServiceFacade
	EventBus    cqrsInterfaces.EventBus
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		LinkService: deps.LinkService,
		EventBus:    deps.EventBus,
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

		query := linkQueries.GetByHashQuery{
			Params: &linkPayloads.GetByHashParams{
				Hash: hash,
			},
		}

		result, err := handler.LinkService.Queries.GetByHash(ctx, query)

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
		filter := linkPayloads.GetAllParams{}
		err := req.ParseQuery(r, &filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := linkQueries.GetAllQuery{
			Params: &filter,
		}

		result, err := handler.LinkService.Queries.GetAll(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, linkResponses.NewPublicLinkList(result.Items, &result.Count), http.StatusOK)
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

		query := linkQueries.GetByIdQuery{
			Params: &linkPayloads.GetByIDParams{
				Id: *id,
			},
		}

		result, err := handler.LinkService.Queries.GetById(ctx, query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusOK)
	}
}

func (handler *Handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := req.HandleBody[linkPayloads.CreatePayload](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := linkCommands.CreateCommand{
			Payload: payload,
		}

		result, err := handler.LinkService.Commands.Create(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusCreated)
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

		payload, err := req.HandleBody[linkPayloads.UpdateRequest](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := linkCommands.UpdateCommand{
			Payload: &linkPayloads.UpdatePayload{
				Id:   *id,
				Url:  payload.Url,
				Hash: payload.Hash,
			},
		}

		result, err := handler.LinkService.Commands.Update(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusOK)
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

		cmd := linkCommands.DeleteCommand{
			Payload: &linkPayloads.GetByIDParams{
				Id: *id,
			},
		}

		_, err = handler.LinkService.Commands.Delete(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
