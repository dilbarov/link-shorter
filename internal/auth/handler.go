package auth

import (
	"link-shorter/configs"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"log"
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
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[LoginRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		res.Json(w, &LoginResponse{
			Token: "123",
		}, 200)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("register")
	}
}
