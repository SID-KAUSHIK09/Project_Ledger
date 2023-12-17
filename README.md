# Zopsmart's Project Ledger

Zopsmart's Project Ledger is a CRUD API built in GoFr, an opinionated microservice development framework based on the Go language. This API is designed to manage projects at Zopsmart, handling project names, types, and status, providing essential CRUD operations.

## API Overview

- **C (Create):** Create a project by submitting a POST request with the project's name, type, and status.
- **R (Read):** Retrieve project details using GET requests.
- **U (Update):** Update a project's name, type, or status through PUT requests.
- **D (Delete):** Delete any project by specifying its ID in a DELETE request.

All API calls can be conveniently performed using Postman.

## Constraints

Constraints have been applied to project types and statuses:

- Project types: "ecommerce", "logistics", "retail", "supplychain", or "others".
- Status: "inprocess" or "completed".

Testing for CRUD functionalities is included in the `main_test.go` file.

## Sequence Diagram
![Blank diagram (1)](https://github.com/SID-KAUSHIK09/Project_Ledger/assets/108971849/e3cca5b4-7b11-40e6-b18f-37adc282bef7)


## Getting Started

Follow these steps to run the code:

1. Install GO, Docker, Postman, and VsCode.
2. Initialize the Go module using the command: `go mod init github.com/example`.
3. Add the GoFr package to the project: `go get gofr.dev`.
4. Download and sync required modules: `go mod tidy`.
5. Create a "configs" folder in the project and add a `.env` file.
6. Connect to the MYSQL database using Docker:
    ```bash
    docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
    ```
7. Update .env file to as mentioned in code.
8. Create the table in the database:
    ```bash
    docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE projects (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, ptype VARCHAR(255) NOT NULL, status VARCHAR(255) NOT NULL);"
    ```
9. Run the server: `go run main.go`.
10. Open Postman and use the following requests for CRUD operations.

### Postman Requests

- **GET:** `http://localhost:9000/project` - Retrieve all project details.
  
- **POST:** `http://localhost:9000/project` - Create a new project. Example JSON:
  ```json
  {
    "name": "Project5",
    "ptype": "ecommerce",
    "status": "inprocess"
  }
  
- **DELETE:** `http://localhost:9000/project/<project_id>` - Delete a project.
Example:` http://localhost:9000/project/8`.

- **PUT:** `http://localhost:9000/project/<project_id>` - Update a project.
Example JSON:
`{
  "name": "Project3",
  "ptype": "ecommerce",
  "status": "completed"
}`

- **Postman Collection:** `https://drive.google.com/file/d/1S8b89qsJqYxraR6D0_mKxAbnwMok18l9/view?usp=drive_link`

## Testing
Install the testify/assert package: `go get github.com/stretchr/testify/assert`.

Run tests: `go test -v`.

## References
`https://gofr.dev/docs`

## Authors

- [Siddharth Kaushik](https://github.com/SID-KAUSHIK09)

