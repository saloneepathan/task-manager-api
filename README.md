# Task Manager API (Go)

A simple, lightweight, and production-ready **RESTful API** built using **Go** and **Gorilla Mux**.  
It allows you to manage tasks with CRUD operations, in-memory storage, and optional Docker support.  

This project is great for showcasing **Go development**, **API design**, and **containerization** skills.

---

## Features
- **CRUD Operations** for tasks  
- **In-Memory Storage** (easily replaceable with a database)  
- **Docker Support** for easy deployment  
- **Makefile** commands for running, testing, and building  
- Clean project structure and code organization  

---

## Tech Stack
- **Language:** Go 1.24+
- **Framework:** Gorilla Mux (Routing)
- **UUID Generation:** Google UUID
- **Containerization:** Docker
- **Build Automation:** Makefile

---

## Project Structure
```

task-manager-api/
│── main.go                # Entry point
│── go.mod, go.sum         # Dependencies
│── handlers/              # API endpoint handlers
│── models/                # Data models
│── routes/                # Route registration
│── storage/               # In-memory storage
│── Makefile               # Run & Docker commands
│── Dockerfile             # Containerization setup
│── README.md              # Documentation

````

---

## Getting Started

### **1. Clone the Repository**
```bash
git clone https://github.com/saloneepathan/task-manager-api.git
cd task-manager-api
````

### **2. Install Dependencies**

```bash
go mod tidy
```

---

## Running the API

### **Option 1: Local Development**

```bash
make run
```

Your API will run at: **[http://localhost:8080](http://localhost:8080)**

---

### **Option 2: Using Docker**

Build and run inside Docker:

```bash
make docker-build
make docker-run
```

Access it at: **[http://localhost:8080](http://localhost:8080)**

---

## API Endpoints

| Method   | Endpoint      | Description       |
| -------- | ------------- | ----------------- |
| `GET`    | `/tasks`      | Get all tasks     |
| `GET`    | `/tasks/{id}` | Get a task by ID  |
| `POST`   | `/tasks`      | Create a new task |
| `PUT`    | `/tasks/{id}` | Update a task     |
| `DELETE` | `/tasks/{id}` | Delete a task     |

---

## Example Usage (cURL)

### **Create a Task**

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title": "Learn Go"}'
```

### **Get All Tasks**

```bash
curl http://localhost:8080/tasks
```

### **Update a Task**

```bash
curl -X PUT http://localhost:8080/tasks/<task_id> \
-H "Content-Type: application/json" \
-d '{"title": "Learn GoLang", "status": "completed"}'
```

### **Delete a Task**

```bash
curl -X DELETE http://localhost:8080/tasks/<task_id>
```

---

## Makefile Commands

| Command             | Description          |
| ------------------- | -------------------- |
| `make run`          | Run API locally      |
| `make test`         | Run tests (if added) |
| `make docker-build` | Build Docker image   |
| `make docker-run`   | Run Docker container |

---

## License

This project is licensed under the **MIT License**.
You are free to use, modify, and distribute this software as per the license terms.
