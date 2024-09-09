package main

import (
	"log"
	"net/http"
	"tonyadjei/go-blog.git/pkg/config"
	"tonyadjei/go-blog.git/pkg/handlers"
)

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr: ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}