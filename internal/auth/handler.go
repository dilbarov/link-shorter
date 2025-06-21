package auth

import (
	"link-shorter/configs"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	AuthService *ServiceFacade
}

type Handler struct {
	*configs.Config
	AuthService *ServiceFacade
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := req.HandleBody[LoginRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		cmd := LoginCommand{Payload: LoginRequest{
			Email:    payload.Email,
			Password: payload.Password,
		}}

		token, err := h.AuthService.Login(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		res.Json(w, &LoginResponse{
			Token: token,
		}, 200)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload, err := req.HandleBody[RegisterRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		cmd := RegisterCommand{Payload: RegisterRequest{
			Email:    payload.Email,
			Password: payload.Password,
			Name:     payload.Name,
		}}

		token, err := h.AuthService.Register(ctx, cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, &RegisterResponse{
			Token: token,
		}, 200)
	}
}
