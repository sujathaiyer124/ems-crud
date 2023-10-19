package main

import (
	"fmt"
	"log"
	"net/http"
	"task2/router"
)

func main() {
	fmt.Println("Welcome to employee management system")
	r := router.Router()
	fmt.Println("Server  is getting started ....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
