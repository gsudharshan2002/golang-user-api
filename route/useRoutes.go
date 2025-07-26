package routes

import (
	"user-api/controller"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
