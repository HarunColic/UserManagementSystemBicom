package routes

import (
	"UserManagementSystem/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoleRoutes = func(router *mux.Router) {

	router.HandleFunc("/roles/", controllers.CreateRole).Methods("POST")
	router.HandleFunc("/roles/", controllers.GetRoles).Methods("GET")
	router.HandleFunc("/roles/{roleId}", controllers.GetRole).Methods("GET")
	router.HandleFunc("/roles/{roleId}", controllers.UpdateRole).Methods("PUT")
	router.HandleFunc("/roles/{roleId}", controllers.DeleteRole).Methods("DELETE")

}
