package main

import (
	"lagom/db"
	"lagom/routing"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type Router struct {
	*chi.Mux
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db.InitDB()

	router := routing.Router(public())

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
