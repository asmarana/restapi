package routes

import (
	"github.com/gorilla/mux"
)

var BasicCrudRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/students", controllers.getStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.getStudent).Methods("GET")
	router.HandleFunc("/api/students/create", controllers.createStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.updateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.deleteStudent).Methods("DELETE")
}
