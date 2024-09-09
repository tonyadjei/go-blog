package main

import (
	"log"
	"net/http"
	"time"
	"tonyadjei/go-blog.git/pkg/config"
	"tonyadjei/go-blog.git/pkg/handlers"

	"github.com/alexedwards/scs/v2"
)


var sessionManager *scs.SessionManager
var app config.AppConfig


func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	
	app.Session = sessionManager


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