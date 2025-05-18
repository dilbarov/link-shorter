package link

import (
	"link-shorter/configs"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
}

type Handler struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}

	router.HandleFunc("GET /link/", handler.getAll())
	router.HandleFunc("GET /link/{id}", handler.getById())
	router.HandleFunc("POST /link/{id}", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}

func (handler *Handler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}

func (handler *Handler) getById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[CreateLinkRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}

func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[UpdateLinkRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}

func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res.Json(w, &Link{
			Url:  "",
			Hash: "",
		}, 200)
	}
}
