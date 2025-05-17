package auth

import (
	"encoding/json"
	"link-shorter/configs"
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
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("auth")
		res := LoginResponse{
			Token: "123",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("register")
	}
}
