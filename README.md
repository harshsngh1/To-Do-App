# TODO App API Service

This is a Go API service for a TODO application.

## Introduction

This project implements a RESTful API service for managing tasks in a TODO application. It allows users to create new tasks, edit existing tasks, list all tasks, filter tasks based on criteria such as due date and priority, and mark tasks as done or canceled.

## Technologies Used

- Language: Go
- Framework: Echo (chosen for its simplicity, performance, and ease of use)
- Database: SQLite (intended for future use as a lightweight and easy-to-setup solution)
- Test Framework: Go testing (standard library for writing tests in Go)

## Project Structure

The project is structured as follows:

todo-app/
│
├── handlers/
│ ├── createTask_handler.go
│ ├── editTask_handler.go
│ ├── listTasks_handler.go
│ └── markTask_handler.go
│
├── models/
│ └── task.go
│
├── routes/
│ └── router.go
│
├── tests/
│ └── to-do_test.go
│
├── main.go
├── go.mod
└── README.md


- `handlers/`: Contains the request handlers for each API endpoint.
- `models/`: Contains the data models used in the application.
- `routes/`: Defines the API routes and links them to corresponding handlers.
- `tests/`: Contains unit tests for the application.
- `main.go`: Entry point of the application.
- `go.mod`: Go module file listing project dependencies.
- `README.md`: Provides an overview of the project and instructions on how to run it.

## Installation

1. Clone the repository:

https://github.com/harshsngh1/To-Do-App


2. Navigate to the project directory:

cd todo-app


3. Run the following command to start the server:

go run main.go


## API Endpoints

- `POST /tasks`: Create a new task
- `GET /tasks`: List all tasks
- `PUT /tasks/:id`: Edit a task by ID
- `PUT /tasks/:id/status`: Mark a task as done or canceled

## Database

While the current implementation uses in-memory maps for storing data, the design is intended to support a database-backed storage solution for scalability and persistence. SQLite has been chosen as the database technology due to its lightweight nature and ease of setup.

## Testing

To run the unit tests, use the following command:

go test ./tests


## Running via Docker

To run the application using Docker, execute the following commands:

1. Build the Docker image:

docker build -t todo-app .

2. Run the Docker container:

docker run -d -p 8080:8080 --name todo-app todo-app

The application will start inside a Docker container and be accessible at http://localhost:8080.
