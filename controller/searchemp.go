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

func SearchEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Id to search:")
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
			salary, err := strconv.ParseFloat(record[7], 64)
			if err != nil {
				log.Printf("Error converting salary to integer: %s", err.Error())
				continue
			}
			employees := models.Employee{
				ID:        id,
				FirstName: record[1],
				LastName:  record[2],
				Email:     record[3],
				Password:  record[4],
				PhoneNo:   record[5],
				Role:      record[6],
				Salary:    salary,
			}
			fmt.Fprintln(w,"Speccific employee deatils are:")
			jsonData, err := json.Marshal(employees)
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
		log.Println("Id not found")
		json.NewEncoder(w).Encode("Id not found")
	}
}
