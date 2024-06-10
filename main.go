package main

import (
	"fmt"
	"log"
	"net/http"

	"go_crud_service/handlers"
	"go_crud_service/repository"
	"go_crud_service/services"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewEmployeeRepository()
	service := services.NewEmployeeService(repo)
	handler := handlers.NewEmployeeHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/employees", handler.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employees", handler.ListEmployees).Methods("GET")

	fmt.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
