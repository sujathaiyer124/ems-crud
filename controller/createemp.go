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
	"task2/models"
	"time"
)

func CreateEmployees(w http.ResponseWriter, r *http.Request, employee []models.Employee) {
	fmt.Println("Create employee")
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
	//var course struct_emp.Employee
	var createemp []models.Employee
	err = json.NewDecoder(r.Body).Decode(&createemp)
	if err != nil {
		json.NewEncoder(w).Encode("No data inside Json")
		return
	}
	//fmt.Printf("Decoded data: %+v\n", createemp)
	createemp = append(employee, createemp...)
	//fmt.Printf("Decoded data: %+v\n", createemp)

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
	json.NewEncoder(w).Encode(createemp)
	json.NewEncoder(w).Encode("Data added successfully")

}
