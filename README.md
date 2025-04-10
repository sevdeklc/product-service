# Product Service

##### Goal

Writing a microservice for performing product operations

## Table of Contents
1. [Tech Stack](#tech-stack)
2. [Requirements](#requirements)
3. [Database Configuration](#database-configuration)
4. [Running the Application](#running-the-application)
5. [Run Swagger UI Documentation](#run-swagger-ui-documentation)


## Tech Stack

- Go (Golang) 1.24.1
- Gin Gonic
- GORM
- MySQL
- Swagger (swaggo)

## Requirements

For building and running the application, the following are required:

- Go 1.24.1 or higher
- MySQL
- Swagger CLI (optional, for generating API docs: swag init)

Note: Make sure the .env file is located in the root directory and contains all the required environment variables.

## Database Configuration

To run the Product Service, you need to create a MySQL database schema named `product-service` (or your preferred name).

Make sure to define the database connection credentials in the .env file located in the root directory of the project:

```
DB_USER=<your_username>
DB_PASSWORD=<your_password>
DB_NAME=<your_db>
DB_HOST=<host>
DB_PORT=<port>
```

The application will use these environment variables to establish a connection via GORM.

## Running the application

The Product Service runs on port 8080 by default and can be started with the following command:

```shell
go run main.go
```

## Run OpenAPI - Swagger UI Documentation

After running the application, you can access the Swagger UI for the Product Service at the following URL:

👉 http://localhost:8081/swagger/index.html

Note: Make sure you have already generated the Swagger docs using the following command (if not done yet):
```shell
swag init
```