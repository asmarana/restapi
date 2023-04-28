package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/restapi/models"
	"github.com/restapi/utils"

	"github.com/gorilla/mux"
)

var students []models.Student

// Get All Students
func getStudents(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// Get single book
func getStudent(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	// Gets params
	// Loop through books and find one with the id from the params
	for _, item := range students {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Student{})
}
func createStudent(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student models.Student
	utils.ParseBody(router, student)

	_ = json.NewDecoder(router.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(10000000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

// Update books
func updateStudent(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var student models.Student
			_ = json.NewDecoder(router.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
}

// Delete book
func deleteStudent(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}
