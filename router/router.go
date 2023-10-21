package router

import (
	"net/http"
	createemployee "task2/controller/createemp"
	deleteemployee "task2/controller/deleteemp"
	reademployee "task2/controller/reademp"
	"task2/controller/searchbyemail"
	search "task2/controller/searchbyfirstname"
	searchemployee "task2/controller/searchbyid"
	searchbylastname "task2/controller/searchbylname"
	"task2/controller/searchbyrole"
	updateemployee "task2/controller/updateemp"
	"task2/models"

	"github.com/gorilla/mux"
)

var employees []models.Employee

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		createemployee.CreateEmployees(w, r, employees)
	}).Methods("POST")

	r.HandleFunc("/employees", reademployee.ReadEmployee).Methods("GET")
	r.HandleFunc("/employees/{id}", searchemployee.SearchEmployee).Methods("GET")
	r.HandleFunc("/employees/{id}", updateemployee.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteemployee.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employees-firstname", search.SearchEmployeeByFirstName).Methods("GET").Queries("firstname", "{firstname}")
	r.HandleFunc("/employees-lastname", searchbylastname.SearchEmployeeByLastName).Methods("GET").Queries("lastname", "{lastname}")
	r.HandleFunc("/employees-email", searchbyemail.SearchEmployeeByEmail).Methods("GET").Queries("email", "{email}")
	r.HandleFunc("/employees-role", searchbyrole.SearchEmployeeByRole).Methods("GET").Queries("role", "{role}")

	return r
}
