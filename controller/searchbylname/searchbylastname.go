package searchbylastname

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

func SearchEmployeeByLastName(w http.ResponseWriter, r *http.Request) {
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
		log.Fatal("Data not read", err)
	}
	defer f.Close()
	var found bool
	
	params := mux.Vars(r)
	lastname := params["lastname"]
	var foundRecord models.Employee
	for i := 1; i < len(data); i++ {
		record := data[i]
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}

		if record[2] == lastname {
			found = true
			log.Println(record[2])
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
			jsonData, err := json.Marshal(foundRecord)
			if err != nil {
				log.Fatalf("Error encoding JSON: %s", err.Error())
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Record not found")
		json.NewEncoder(w).Encode("Record not found")
	}
}
