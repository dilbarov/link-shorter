package auth

import (
	"link-shorter/configs"
	authPayloads "link-shorter/internal/auth/payloads"
	authServices "link-shorter/internal/auth/services"
	authCommands "link-shorter/internal/auth/services/commands"
	"link-shorter/pkg/req"
	"link-shorter/pkg/res"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	AuthService *authServices.ServiceFacade
}

type Handler struct {
	*configs.Config
	AuthService *authServices.ServiceFacade
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
		_, err := req.HandleBody[authPayloads.LoginRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		res.Json(w, &authPayloads.LoginResponse{
			Token: "123",
		}, 200)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[authPayloads.RegisterRequest](r.Body)

		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		cmd := authCommands.RegisterCommand{Payload: authPayloads.RegisterRequest{
			Email:    payload.Email,
			Password: payload.Password,
			Name:     payload.Name,
		}}

		token, err := h.AuthService.Commands.Register.Execute(cmd)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		res.Json(w, &authPayloads.RegisterResponse{
			Token: token,
		}, 200)
	}
}
