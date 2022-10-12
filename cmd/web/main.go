package main

import (
	"fmt"
	"log"
	"net/http"
	"webHello/pkg/config"
	"webHello/pkg/handlers"
	"webHello/pkg/render"
)

const portNumber = ":8888"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.Env = "prod"
	app.UseCache = !(app.Env == "dev")

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server start in : http://localhost%s", portNumber))
	fmt.Println(fmt.Sprintf("Used Cache template : %v", app.UseCache))
	fmt.Println(fmt.Sprintf("Envirenement : %s", app.Env))
	err = http.ListenAndServe(portNumber, nil)

	if err != nil {
		log.Println("Server error", err)
	}

}
