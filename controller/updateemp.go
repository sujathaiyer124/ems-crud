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
			// headers := []string{"ID", "FirstName", "LastName", "Email", "Password", "Phoneno", "Role", "Salary"}
			// writer.Write(headers)
			var employee []models.Employee

			err = json.NewDecoder(r.Body).Decode(&employee)
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
			json.NewEncoder(w).Encode("Employee updated successfully")
			json.NewEncoder(w).Encode(data)
			return

		}

	}
	if !found {
		log.Println("ID not found")
		json.NewEncoder(w).Encode("Employee not found")
	}
}
