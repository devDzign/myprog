package main

import (
	"fmt"
	http "net/http"
)

func WriteToConsole(next http.Handler) http.Handler {

	fun := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi the page")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fun)
}
