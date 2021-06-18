package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers set the port, handler and start server
func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8080"
	}

	handler := cors.AllowAll().Handler(router) // give access permissions from everywhere

	log.Fatal(http.ListenAndServe(":"+PORT, handler)) 
}