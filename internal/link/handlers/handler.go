package handlers

import (
	"link-shorter/configs"
	linkModels "link-shorter/internal/link/models"
	linkPayloads "link-shorter/internal/link/payloads"
	linkResponses "link-shorter/internal/link/responses"
	linkServices "link-shorter/internal/link/services"
	linkCommands "link-shorter/internal/link/services/commands"
	"link-shorter/internal/link/services/queries"
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

	router.HandleFunc("GET /link/", handler.getAll())
	router.HandleFunc("GET /link/{id}", handler.getById())
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		query := queries.GetByHashQuery{
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

		res.Json(w, &linkModels.Model{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}

func (handler *Handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		query := queries.GetByIdQuery{
			Params: &linkPayloads.GetByIDParams{
				Id: id,
			},
		}

		result, err := handler.LinkService.Queries.GetById.Execute(query)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, linkResponses.NewPublicResponse(result), http.StatusOK)
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

		res.Json(w, linkResponses.NewPublicResponse(result), http.StatusCreated)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ExtractUUID(r, "id")

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

		res.Json(w, linkResponses.NewPublicResponse(result), http.StatusOK)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := req.ExtractUUID(r, "id")
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
