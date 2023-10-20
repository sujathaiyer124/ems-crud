package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"task2/models"

	"github.com/gorilla/mux"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the employee data")
	w.Header().Set("Content-Type", "application/json")
	f, err := os.OpenFile("emp.csv", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal("File is not opened")
	}
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Data not read", err)
	}
	f.Close()

	e := os.Remove("emp.csv")
	if e != nil {
		log.Fatal(e)
	}
	var found bool
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Cannot convert from string to int hhhh %s", err.Error())
	}
	for index, record := range data {
		recordIdStr := record[0]
		if recordIdStr == "" {
			log.Println("Skipping empty string at index", index)
			continue
		}
		if recordIdStr == "ID" {
			log.Println("Skipping header row with value 'ID' at index", index)
			continue
		}
		recordId, err := strconv.Atoi(recordIdStr)

		if err != nil {
			log.Println("Error converting to integer at index", index)
			log.Println("Value causing the error:", recordIdStr)
			log.Println("Error details:", err)
		}
		if recordId == id {
			found = true
			file, err := os.Create("emp.csv")
			if err != nil {
				log.Fatal("File is not been created")
			}
			defer file.Close()
			writer := csv.NewWriter(file)
			defer writer.Flush()

			var employee []models.Employee

			err = json.NewDecoder(r.Body).Decode(&employee)
			for _, e := range employee {
				if e.Salary <= 0 {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(w, "Salary must be greater than 0")
					return
				}
				if err := models.ValidateEmployee(e); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(err.Error()))
					return
				}
				if !models.CustomPasswordValidation(e.Password) {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(w, "Invalid password. Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character.")
					return
				}
				if e.Role != "admin" && e.Role != "user" {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(w, "Role must be admin or user")
					return
				}

			}
			if err != nil {
				log.Println("error is", err)
				return
			}
			for _, emp := range employee {
				data[index] = []string{
					strconv.Itoa(id),
					emp.FirstName,
					emp.LastName,
					emp.Email,
					emp.Password,
					emp.PhoneNo,
					emp.Role,
					strconv.FormatFloat(emp.Salary, 'f', 2, 64),
				}
				employee = append(employee, emp)
			}
			if index < len(employee) {
				if index < len(employee)-1 {
					data = append(data[:index], data[index+1:]...)
				} else {
					employee = employee[:index]
				}
			} else {
				log.Println("Index is out of bounds")
			}

			writer.WriteAll(data)
			log.Println("Employee after append:", employee)
			//fmt.Fprintln(w, "The details of the employee after update are:")
			jsonData, err := json.Marshal(employee)
			if err != nil {
				log.Fatalf("Error encoding JSON: %s", err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			return

		}

	}
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ID not found")
		json.NewEncoder(w).Encode("Employee not found")
	}
}
