package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"go_crud_service/models"
	"go_crud_service/repository"
	"go_crud_service/services"

	"github.com/gorilla/mux"
)

func setup() (*mux.Router, *repository.EmployeeRepository) {
	repo := repository.NewEmployeeRepository()
	service := services.NewEmployeeService(repo)
	handler := NewEmployeeHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/employees", handler.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id:[0-9]+}", handler.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employees", handler.ListEmployees).Methods("GET")
	return r, repo
}

func TestCreateEmployeeHandler(t *testing.T) {
	router, _ := setup()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	body, _ := json.Marshal(employee)

	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdEmployee models.Employee
	json.Unmarshal(rr.Body.Bytes(), &createdEmployee)
	if createdEmployee.Name != employee.Name {
		t.Errorf("expected employee Name to be %s, got %s", employee.Name, createdEmployee.Name)
	}
}

func TestGetEmployeeByIDHandler(t *testing.T) {
	router, repo := setup()
	employee := repo.CreateEmployee(models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000})

	req, _ := http.NewRequest("GET", "/employees/"+strconv.Itoa(employee.ID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var fetchedEmployee models.Employee
	json.Unmarshal(rr.Body.Bytes(), &fetchedEmployee)
	if fetchedEmployee.Name != employee.Name {
		t.Errorf("expected employee Name to be %s, got %s", employee.Name, fetchedEmployee.Name)
	}
}

func TestUpdateEmployeeHandler(t *testing.T) {
	router, repo := setup()
	employee := repo.CreateEmployee(models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000})
	updatedEmployee := models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000}
	body, _ := json.Marshal(updatedEmployee)

	req, _ := http.NewRequest("PUT", "/employees/"+strconv.Itoa(employee.ID), bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	fetchedEmployee, _ := repo.GetEmployeeByID(employee.ID)
	if fetchedEmployee.Name != updatedEmployee.Name {
		t.Errorf("expected employee Name to be %s, got %s", updatedEmployee.Name, fetchedEmployee.Name)
	}
}

func TestDeleteEmployeeHandler(t *testing.T) {
	router, repo := setup()
	employee := repo.CreateEmployee(models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000})

	req, _ := http.NewRequest("DELETE", "/employees/"+strconv.Itoa(employee.ID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	_, err := repo.GetEmployeeByID(employee.ID)
	if err == nil {
		t.Errorf("expected error when fetching deleted employee, got nil")
	}
}

func TestListEmployeesHandler(t *testing.T) {
	router, repo := setup()
	repo.CreateEmployee(models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000})
	repo.CreateEmployee(models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000})

	req, _ := http.NewRequest("GET", "/employees?offset=0&limit=1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var employees []models.Employee
	json.Unmarshal(rr.Body.Bytes(), &employees)
	if len(employees) != 1 {
		t.Errorf("expected 1 employee, got %d", len(employees))
	}
}
