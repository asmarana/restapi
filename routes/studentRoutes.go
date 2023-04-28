package routes

import (
	"github.com/gorilla/mux"
	"github.com/restapi/controllers"
)

var BasicCrudRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/api/students/create", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.DeleteStudent).Methods("DELETE")
}
