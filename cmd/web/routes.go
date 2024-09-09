package main

import (
	"net/http"
	"tonyadjei/go-blog.git/pkg/config"
	"tonyadjei/go-blog.git/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}