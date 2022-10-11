package handlers

import (
	"net/http"
	"webHello/pkg/render"
)

func Home(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
