package main

import (
	"fmt"
	"log"
	"net/http"
	"webHello/pkg/handlers"
)

const portNumber = ":8888"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server start in : http://localhost%s", portNumber))
	err := http.ListenAndServe(portNumber, nil)

	if err != nil {
		log.Println("Server error", err)
	}

}
