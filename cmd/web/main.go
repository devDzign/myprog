package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"webHello/pkg/config"
	"webHello/pkg/handlers"
	"webHello/pkg/render"
)

const portNumber = ":8888"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	isProduction := !(app.Env == "dev")
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = isProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.Env = "prod"
	app.UseCache = isProduction
	app.IsProduction = isProduction

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Server start in : http://localhost%s", portNumber))
	fmt.Println(fmt.Sprintf("Used Cache template : %v", app.UseCache))
	fmt.Println(fmt.Sprintf("Envirenement : %s", app.Env))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
