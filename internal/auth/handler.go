package auth

import (
	"encoding/json"
	"link-shorter/configs"
	"link-shorter/pkg/res"
	"log"
	"net/http"
	"net/mail"
	"regexp"
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
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 400)
			return
		}

		log.Printf("%s", payload.Password)

		if payload.Email == "" {
			res.Json(w, "email is empty", 400)
			return
		}

		match, _ := regexp.MatchString(`[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}`, payload.Email)

		if !match {
			res.Json(w, "email is invalid", 400)
			return
		}

		_, err = mail.ParseAddress(payload.Email)

		if err != nil {
			res.Json(w, "email is invalid", 400)
		}

		if payload.Password == "" {
			res.Json(w, "password is empty", 400)
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
