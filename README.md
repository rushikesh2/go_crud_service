# go_crud_service

Assignment
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
   ![POST](https://user-images.githubusercontent.com/55944826/151755172-4d1af9a8-33ad-4824-a852-b3d9bcd5a98a.png)

2. GET api to all users.
   ![GET](https://user-images.githubusercontent.com/55944826/151755179-f39ce830-e0a9-4041-ad9b-a70dc3ba6add.png)

3. GET by id

4. PUT api to create or update the employee

5. DELETE api to delete the employee
