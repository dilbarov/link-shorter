package auth

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("auth")
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
	}
}

func NewHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}
