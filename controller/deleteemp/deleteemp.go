package deleteemployee

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

	"github.com/gorilla/mux"
)


func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
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
			data = append(data[:index], data[index+1:]...)
			rand.Seed(time.Now().UnixNano())
			file, err := os.Create("emp.csv")
			if err != nil {
				log.Fatal("File is not been created")
			}
			defer file.Close()
			writer := csv.NewWriter(file)
			defer writer.Flush()
			writer.WriteAll(data)
			log.Println("Employee deleted successfully")
			var result []models.Employee
			for i := 1; i < len(data); i++ {
				record := data[i]
				if len(record) != 8 {
					continue
				}
				id, err := strconv.Atoi(record[0])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				salary, err := strconv.ParseFloat(record[7], 64)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				emp := models.Employee{
					ID:        id,
					FirstName: record[1],
					LastName:  record[2],
					Email:     record[3],
					Password:  record[4],
					PhoneNo:   record[5],
					Role:      record[6],
					Salary:    salary,
				}
				result = append(result, emp)
			}

			jsonData, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
		}
	}
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Id not found")
		json.NewEncoder(w).Encode("employee not found")
	}
}
