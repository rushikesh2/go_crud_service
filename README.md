# go_crud_service

# Assignment
You are tasked with creating a Go application that manages a simple employee database or in-memory store. Additionally, you need to implement a RESTful API with pagination for listing employee records.

1. Employee Struct:

- Define a struct named `Employee` with the following fields:
- `ID` (int): Unique identifier for the employee.
- `Name` (string): Name of the employee.
- `Position` (string): Position/title of the employee.
- `Salary` (float64): Salary of the employee.

2. CRUD Operations:

- Implement functions/methods to perform CRUD operations on the employee database or in-memory store:
- `CreateEmployee`: Adds a new employee to the database or store.
- `GetEmployeeByID`: Retrieves an employee from the database or store by ID.
- `UpdateEmployee`: Updates the details of an existing employee.
- `DeleteEmployee`: Deletes an employee from the database or store by ID

3. Concurrency:

- Ensure that the application is safe for concurrent use by using appropriate synchronization mechanisms.

4. Testing:

- Write unit tests to cover the CRUD operations and ensure the correctness of the implementation.

5. RESTful API with Pagination:

- Implement a RESTful API for listing employee records with pagination.
- The API should provide endpoints for listing employees with support for pagination.
- Each page should contain a configurable number of records.
- Implement proper error handling and response formatting for the API endpoints.

## APIs:

1. POST api to create employee.
![POST request](https://github.com/rushikesh2/go_crud_service/assets/49398834/660d53f3-c8da-4a98-9a01-50e7432990f1)

2. GET api to all employees.
![GET All employee](https://github.com/rushikesh2/go_crud_service/assets/49398834/5be3716e-7061-4764-bc86-61c1dfd07f7c)

3. GET by id
![GET a employee](https://github.com/rushikesh2/go_crud_service/assets/49398834/7ff64331-43fa-45c4-acee-09b3d8153687)

4. PUT api to create or update the employee
![PUT Request](https://github.com/rushikesh2/go_crud_service/assets/49398834/2efedd11-2644-4a6b-99d1-8511ef970549)

6. DELETE api to delete the employee
![Delete Request](https://github.com/rushikesh2/go_crud_service/assets/49398834/6712381a-2348-43b5-9fd6-e72b36b1bb6f)

