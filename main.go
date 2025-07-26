package main

import (
	"log"
	"net/http"

	"user-api/config"
	"user-api/route"

	"github.com/gorilla/mux"
)

func main() {
	
	config.ConnectDB()

	
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
