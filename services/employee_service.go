package services

import (
	"go_crud_service/models"
	"go_crud_service/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee models.Employee) models.Employee {
	return s.repo.CreateEmployee(employee)
}

func (s *EmployeeService) GetEmployeeByID(id int) (models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *EmployeeService) UpdateEmployee(id int, updatedEmployee models.Employee) error {
	return s.repo.UpdateEmployee(id, updatedEmployee)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repo.DeleteEmployee(id)
}

func (s *EmployeeService) ListEmployees(offset, limit int) []models.Employee {
	return s.repo.ListEmployees(offset, limit)
}
