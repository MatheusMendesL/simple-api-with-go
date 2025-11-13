package main

import (
	api "_039_projeto3/routes"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := runServer(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("All systems offline")
}

func runServer() error {
	r := chi.NewMux()

	r.Mount("/", api.ControlRoutes())
	fmt.Println("Server running!")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return err
	}
	slog.Info("The API is running!")
	return nil
}
