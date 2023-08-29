# Contact Pagination Project

This project demonstrates contact pagination using GORM, Gin, Testify, OpenAPI Swagger, and MySQL.

## How to run the project?

We can run this Go boilerplate project with following steps:

- Move to your workspace: `cd your-workspace`.
- Clone this project into your workspace: `git clone https://github.com/prayogatriady/contact-pagination`.
- Move to the project root directory: `cd contact-pagination`.
- Create a file `.env` similar to existing `.env` file at the root directory for your environment variables.

```bash
DB_USER=your-db-user
DB_PASSWORD=your-db-pass
DB_HOST=your-db-host
DB_PORT=your-db-port
DB_NAME=your-db-name
PORT=your-port

DB_USER_TEST=your-db-user-in-testing-environment
DB_PASSWORD_TEST=your-db-pass-in-testing-environment
DB_HOST_TEST=your-db-host-in-testing-environment
DB_PORT_TEST=your-db-port-in-testing-environment
DB_NAME_TEST=your-db-name-in-testing-environment
PORT_TEST=your-port-in-testing-environment
```

- [Install `go`](https://go.dev/doc/install) if not installed on your machine.
- Create table manually, you can copy-paste the sql syntax from migration folder
- Run `go run cmd/main.go`.
- Example to access the pagination API [http://localhost:8000/api/contacts?limit=2&page=1&sort=id%20asc](http://localhost:8000/api/contacts?limit=2&page=1&sort=id%20asc).

## How to run the tests?

```bash
# Run all tests
go test ./...
```
