package handlers

import (
	"link-shorter/configs"
	linkPayloads "link-shorter/internal/link/payloads"
	linkResponses "link-shorter/internal/link/responses"
	linkServices "link-shorter/internal/link/services"
	linkCommands "link-shorter/internal/link/services/commands"
	linkQueries "link-shorter/internal/link/services/queries"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	LinkService *linkServices.ServiceFacade
}

type Handler struct {
	*configs.Config
	LinkService *linkServices.ServiceFacade
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		LinkService: deps.LinkService,
	}

	router.HandleFunc("GET /links", handler.getAll())
	router.HandleFunc("GET /links/{id}", handler.getById())
	router.HandleFunc("POST /links", handler.Create())
	router.HandleFunc("PATCH /links/{id}", handler.Update())
	router.HandleFunc("DELETE /links/{id}", handler.Delete())
	router.HandleFunc("GET /r/{hash}", handler.GoTo())
}

func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		query := linkQueries.GetByHashQuery{
			Params: &linkPayloads.GetByHashParams{
				Hash: hash,
			},
		}

		result, err := handler.LinkService.Queries.GetByHash.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, result.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := linkPayloads.GetAllParams{}
		err := req.ParseQuery(r, &filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := linkQueries.GetAllQuery{
			Params: &filter,
		}

		links, count, err := handler.LinkService.Queries.GetAll.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, linkResponses.NewPublicLinkList(links, &count), http.StatusOK)
	}
}

func (handler *Handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		query := linkQueries.GetByIdQuery{
			Params: &linkPayloads.GetByIDParams{
				Id: id,
			},
		}

		result, err := handler.LinkService.Queries.GetById.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusOK)
	}
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[linkPayloads.CreatePayload](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := linkCommands.CreateCommand{
			Payload: payload,
		}

		result, err := handler.LinkService.Commands.Create.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusCreated)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		result, err := handler.LinkService.Commands.Update.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, linkResponses.NewPublicLink(result), http.StatusOK)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = handler.LinkService.Commands.Delete.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
