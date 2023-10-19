package controller

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"task2/models"	
)
func ReadEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	f, err := os.Open("emp.csv")
	if err != nil {
		log.Fatal("File is not opened")
	}
	defer f.Close()
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Data not read", err)
	}
	var emp []models.Employee
	for _, record := range data {
		id, err := strconv.Atoi(record[0])
		if err != nil {   //error yaha pe aa raha hai
			log.Printf("Error converting ID to integer: %s", err.Error())
			continue
		}
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
		emp = append(emp, employees)
		log.Println(emp)
	}
	json.NewEncoder(w).Encode("All the Employee details:")
	json.NewEncoder(w).Encode(emp)
}