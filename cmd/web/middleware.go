package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	http "net/http"
)

func WriteToConsole(next http.Handler) http.Handler {

	fun := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi the page")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fun)
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
