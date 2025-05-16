package main

import (
	"fmt"
	"link-shorter/internal/auth"
	"net/http"
)

func main() {
	//conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Starting server on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
