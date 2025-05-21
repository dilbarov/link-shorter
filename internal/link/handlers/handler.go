package handlers

import (
	"link-shorter/configs"
	"link-shorter/internal/link/models"
	"link-shorter/internal/link/payloads"
	services "link-shorter/internal/link/services"
	commands "link-shorter/internal/link/services/commands"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	LinkService *services.ServiceFacade
}

type Handler struct {
	*configs.Config
	LinkService *services.ServiceFacade
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

		res.Json(w, &models.LinkModel{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &models.LinkModel{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}

func (handler *Handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &models.LinkModel{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[payloads.LinkCreateRequest](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := commands.CreateCommand{
			Payload: payload,
		}

		result, err := handler.LinkService.Commands.CreateHandler.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, result, http.StatusCreated)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[payloads.LinkUpdateRequest](r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, &models.LinkModel{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &models.LinkModel{
			Url:  "",
			Hash: "",
		}, http.StatusOK)
	}
}
