package main

import (
	"log"
	"github.com/rodri-jara/golang-app-tw/db"
	"github.com/rodri-jara/golang-app-tw/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No DB connection")
	}

	handlers.Handlers()
}
