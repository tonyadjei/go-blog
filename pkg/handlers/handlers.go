package handlers

import (
	"net/http"
	"tonyadjei/go-blog.git/models"
	"tonyadjei/go-blog.git/pkg/config"
	"tonyadjei/go-blog.git/pkg/render"
)


type Repository struct{
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}


func (m *Repository) HomeHandler(w http.ResponseWriter,
	request *http.Request){
		render.RenderTemplate(w, "home.page.tmpl",
	&models.PageData{})
	}

func (m *Repository) AboutHandler(w http.ResponseWriter,
request *http.Request){

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love talking about ourselves."

	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap})
}