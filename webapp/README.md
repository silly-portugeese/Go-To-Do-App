# To Do Web App

**The webapp should incorporate a REST API that does the following:**
- Read To Do task
- Create To Do task
- Read / list all To Do tasks
- Delete To Do task 
- Complete To Do task

(Data is in memory only and will be lost on switch off)


## Structure
```
├── webapp
    │
    ├── backend
    │   │         
    │   ├── handlers
    │   │   └── api_handlers.go
    │   │ 
    │   ├── models
    │   │   └── models.go
    │   │     
    │   ├── service
    │   │   └── service.go
    │   │ 
    │   ├── storage
    │   │   ├── storage_benchmark_test.go
    │   │   ├── storage_concurrency_test.go
    │   │   ├── storage_test.go
    │   │   └── storage.go
    │   │ 
    │   └── main.go
    │     
    └── frontend
        │ 
        ├── handlers
        │   └── html_handlers.go
        │ 
        ├── models
        │   └── models.go
        │ 
        ├── static
        │   ├── edit.html
        │   └── view.html
        │          
        └── main.go      
```

## API Endpoints

| **Method** | **URL**                  | **Description**                   |
|------------|--------------------------|-----------------------------------|
| GET        | `/api/todos`             | Retrieve all To Do items.         |
| GET        | `/api/todo/{id}`         | Retrieve a To Do item by ID.      |
| POST       | `/api/todo/`             | Create a new To Do item.          |
| PUT        | `/api/todo/update/{id}`  | Update a To Do item by ID.        |
| DELETE     | `/api/todo//delete/{id}` | Delete a To Do item by ID.        |

###	Request Body:

**POST /api/todo/**

```sh
{
  "task": "Learn Go",
}
```

**PUT /api/todo/update/{id}

```sh
{
  "task": "Learn Go",
  "status": "in progress",  // Options: "pending", "in progress", "completed"
}
```

## HTML Interface

HTMX is used to enhance the HTML interface, making it more dynamic. It is included via a CDN for simplicity. For more details on HTMX, refer its [documentation](https://htmx.org/).


| **Method** | **URL**             | **Description**                   |
|------------|---------------------|-----------------------------------|
| GET        | `/todos`            | View all To Do items.             |
| GET        | `/todo/edit/{id}`   | Edit a To Do item by ID.          |


## Starting the Project

**Start the backend server for the web app:**
```sh
./run.sh webapp-backend
```

**Start the frontend server for the web app:**
```sh
./run.sh webapp-frontend
```

**Start both the backend and frontend server for the web app:**
```sh
./run.sh webapp
```

## Running Tests

To ensure the quality and performance of the application, you can run the following tests:

**Unit Tests**
```sh
./run.sh webapp-unit-test
```

**Benchmark Tests**
```sh
./run.sh webapp-benchmark-test
```

**Concurrency Tests**
```sh
./run.sh webapp-concurrency-test
```