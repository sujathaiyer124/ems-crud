package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	//"strings"
	"task2/models"
	"time"
	
)


func CreateEmployees(w http.ResponseWriter, r *http.Request, employee []models.Employee) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var writer *csv.Writer
	var file *os.File
	filename := "emp.csv"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err = os.Create(filename)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		writer = csv.NewWriter(file)
		defer writer.Flush()
		headers := []string{"ID", "FirstName", "LastName", "Email", "Password", "Phoneno", "Role", "Salary"}
		writer.Write(headers)
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()
		if err != nil {
			log.Fatal("Cannot open the file", err)
		}
		writer = csv.NewWriter(file)
		defer writer.Flush()
	}
	var err error
	var createemp []models.Employee
	err = json.NewDecoder(r.Body).Decode(&createemp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}

	for _, emp := range createemp {
		if emp.Salary <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Salary must be greater than 0")
			return
		}
		if !models.CustomPasswordValidation(emp.Password) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid password. Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character.")
			return
		}
		if err := models.ValidateEmployee(emp); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
        if emp.Role != "admin" && emp.Role != "user" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Role must be admin or user")
			return
		}
	}
	createemp = append(employee, createemp...)
	//fmt.Fprintln(w, "Created employees")
	rand.Seed(time.Now().UnixNano())
	var data []string
	for _, e := range createemp {
		salarystr := strconv.FormatFloat(e.Salary, 'f', 2, 64)
		data = []string{
			strconv.Itoa(rand.Intn(100)),
			e.FirstName,
			e.LastName,
			e.Email,
			e.Password,
			e.PhoneNo,
			e.Role,
			salarystr,
		}
		if err := writer.Write(data); err != nil {
			log.Fatal("Error writing record to CSV: ", err)
		}

	}
	log.Println("Employee Created")
	jsonData, err := json.Marshal(createemp)
	if err != nil {
		log.Fatalf("Error encoding JSON: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
