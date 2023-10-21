package searchbyrole

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

func SearchEmployeeByRole(w http.ResponseWriter, r *http.Request) {
	fmt.Println("name to search:")
	//fmt.Println("Update the employee data")
	w.Header().Set("Content-Type", "application/json")
	f, err := os.Open("emp.csv")
	if err != nil {
		log.Fatal("File is not opened")
	}
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		//fmt.Println(data)
		log.Fatal("Data not read", err)
	}
	defer f.Close()
	var found bool

	params := mux.Vars(r)
	role := params["role"]
	var emp []models.Employee
	var foundRecord models.Employee
	for i := 1; i < len(data); i++ {
		record := data[i]
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}

		if record[6] == role {
			found = true
			//log.Println(record[6])
			salary, err := strconv.ParseFloat(record[7], 64)
			if err != nil {
				log.Printf("Error converting salary to integer: %s", err.Error())
				continue
			}
			foundRecord = models.Employee{
				ID:        id,
				FirstName: record[1],
				LastName:  record[2],
				Email:     record[3],
				Password:  record[4],
				PhoneNo:   record[5],
				Role:      record[6],
				Salary:    salary,
			}
			emp = append(emp, foundRecord)
		}
	}

	if found {
		jsonData, err := json.Marshal(emp)
		if err != nil {
			log.Fatalf("Error encoding JSON: %s", err.Error())
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No records found with the specified role")
	}
}
