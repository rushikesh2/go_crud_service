package repository

import (
	"testing"

	"go_crud_service/models"
)

func TestCreateEmployee(t *testing.T) {
	repo := NewEmployeeRepository()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}

	createdEmployee := repo.CreateEmployee(employee)
	if createdEmployee.ID != 1 {
		t.Errorf("expected employee ID to be 1, got %d", createdEmployee.ID)
	}
	if createdEmployee.Name != employee.Name {
		t.Errorf("expected employee Name to be %s, got %s", employee.Name, createdEmployee.Name)
	}
}

func TestGetEmployeeByID(t *testing.T) {
	repo := NewEmployeeRepository()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := repo.CreateEmployee(employee)

	fetchedEmployee, err := repo.GetEmployeeByID(createdEmployee.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if fetchedEmployee.Name != employee.Name {
		t.Errorf("expected employee Name to be %s, got %s", employee.Name, fetchedEmployee.Name)
	}

	_, err = repo.GetEmployeeByID(999)
	if err == nil {
		t.Errorf("expected error when fetching non-existent employee, got nil")
	}
}

func TestUpdateEmployee(t *testing.T) {
	repo := NewEmployeeRepository()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := repo.CreateEmployee(employee)

	updatedEmployee := models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000}
	err := repo.UpdateEmployee(createdEmployee.ID, updatedEmployee)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	fetchedEmployee, _ := repo.GetEmployeeByID(createdEmployee.ID)
	if fetchedEmployee.Name != updatedEmployee.Name {
		t.Errorf("expected employee Name to be %s, got %s", updatedEmployee.Name, fetchedEmployee.Name)
	}

	err = repo.UpdateEmployee(999, updatedEmployee)
	if err == nil {
		t.Errorf("expected error when updating non-existent employee, got nil")
	}
}

func TestDeleteEmployee(t *testing.T) {
	repo := NewEmployeeRepository()
	employee := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	createdEmployee := repo.CreateEmployee(employee)

	err := repo.DeleteEmployee(createdEmployee.ID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	_, err = repo.GetEmployeeByID(createdEmployee.ID)
	if err == nil {
		t.Errorf("expected error when fetching deleted employee, got nil")
	}

	err = repo.DeleteEmployee(999)
	if err == nil {
		t.Errorf("expected error when deleting non-existent employee, got nil")
	}
}

func TestListEmployees(t *testing.T) {
	repo := NewEmployeeRepository()
	employee1 := models.Employee{Name: "John Doe", Position: "Developer", Salary: 60000}
	employee2 := models.Employee{Name: "Jane Doe", Position: "Manager", Salary: 80000}
	repo.CreateEmployee(employee1)
	repo.CreateEmployee(employee2)

	employees := repo.ListEmployees(0, 1)
	if len(employees) != 1 {
		t.Errorf("expected 1 employee, got %d", len(employees))
	}

	employees = repo.ListEmployees(0, 10)
	if len(employees) != 2 {
		t.Errorf("expected 2 employees, got %d", len(employees))
	}

	employees = repo.ListEmployees(10, 10)
	if len(employees) != 0 {
		t.Errorf("expected 0 employees, got %d", len(employees))
	}
}
