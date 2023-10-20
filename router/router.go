package router

import (
	"net/http"

	"task2/controller"
	"task2/models"

	"github.com/gorilla/mux"
)

var employees []models.Employee

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateEmployees(w, r, employees)
	}).Methods("POST")

	r.HandleFunc("/employees", controller.ReadEmployee).Methods("GET")
	r.HandleFunc("/employees/{id}", controller.SearchEmployee).Methods("GET")
	r.HandleFunc("/employees/{id}", controller.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", controller.DeleteEmployee).Methods("DELETE")
	return r
}
