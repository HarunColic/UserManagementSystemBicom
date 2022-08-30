package routes

import (
	"UserManagementSystem/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {

	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{userId}", controllers.AssignRole).Methods("POST")

}
