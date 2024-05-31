package main

import (
	"log"
	"net/http"

	"github.com/0xMeechie/whatnow/app"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Get("/", app.Home)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln("Unable to Start Server")
	}

}
