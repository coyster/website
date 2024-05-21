package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/go-chi/chi"

	"github.com/coyster/website/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/", handlers.HandleFoo)

	listenAddr := os.Getenv("LISTEN_ADDR")

	fmt.Println("Heelo")
	slog.Info("Listening ")
	http.ListenAndServe(listenAddr, router)

}
