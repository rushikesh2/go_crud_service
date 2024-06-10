package utils

import "go_crud_service/models"

// PaginateEmployees takes a slice of employees and returns a paginated slice
func PaginateEmployees(employees []models.Employee, offset, limit int) []models.Employee {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10 // default limit
	}

	total := len(employees)

	if offset >= total {
		return []models.Employee{}
	}

	end := offset + limit
	if end > total {
		end = total
	}

	return employees[offset:end]
}
