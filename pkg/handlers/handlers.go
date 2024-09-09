package handlers

import (
	"net/http"
	"tonyadjei/go-blog.git/models"
	"tonyadjei/go-blog.git/pkg/render"
)


func HomeHandler(w http.ResponseWriter,
	request *http.Request){
		render.RenderTemplate(w, "home.page.tmpl",
	&models.PageData{})
	}

func AboutHandler(w http.ResponseWriter,
request *http.Request){

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves. We love talking about ourselves."

	render.RenderTemplate(w, "about.page.tmpl", &models.PageData{StrMap: strMap})
}