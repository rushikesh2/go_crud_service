package repository

import (
	"errors"
	"sync"

	"go_crud_service/models"
	"go_crud_service/utils"
)

type EmployeeRepository struct {
	mu        sync.RWMutex
	employees map[int]models.Employee
	nextID    int
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		employees: make(map[int]models.Employee),
		nextID:    1,
	}
}

func (r *EmployeeRepository) CreateEmployee(employee models.Employee) models.Employee {
	r.mu.Lock()
	defer r.mu.Unlock()

	employee.ID = r.nextID
	r.employees[r.nextID] = employee
	r.nextID++
	return employee
}

func (r *EmployeeRepository) GetEmployeeByID(id int) (models.Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	employee, exists := r.employees[id]
	if !exists {
		return models.Employee{}, errors.New("employee not found")
	}
	return employee, nil
}

func (r *EmployeeRepository) UpdateEmployee(id int, updatedEmployee models.Employee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[id]; !exists {
		return errors.New("employee not found")
	}

	updatedEmployee.ID = id
	r.employees[id] = updatedEmployee
	return nil
}

func (r *EmployeeRepository) DeleteEmployee(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[id]; !exists {
		return errors.New("employee not found")
	}

	delete(r.employees, id)
	return nil
}

func (r *EmployeeRepository) ListEmployees(offset, limit int) []models.Employee {
	r.mu.RLock()
	defer r.mu.RUnlock()

	employees := make([]models.Employee, 0, len(r.employees))
	for _, employee := range r.employees {
		employees = append(employees, employee)
	}

	return utils.PaginateEmployees(employees, offset, limit)
}
