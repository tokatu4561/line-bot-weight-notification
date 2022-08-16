package main

import (
	"net/http"
	"tokatu4561/line-bot-weight/handlers"

	"github.com/go-chi/chi/v5"
)

func(app *application) routes() http.Handler  {
	mux := chi.NewRouter()

	// mux.Get("/", handlers.WeightRegist)
	mux.Post("/weight-regist",handlers.WeightRegist)

	return mux
}