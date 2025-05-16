package main

import (
	"errors"
	"fmt"
	"link-shorter/configs"
	"link-shorter/internal/auth"
	"link-shorter/pkg/shutdown"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewHandler(router)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port),
		Handler: router,
	}

	go func() {
		fmt.Printf("Starting server on port %d\n", conf.App.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	shutdown.WaitForShutdown(&server)
}
