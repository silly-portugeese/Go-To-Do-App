# To Do Web App

**The webapp should incorporate a REST API that does the following:**
- Read To Do task
- Create To Do task
- Read / list all To Do tasks
- Delete To Do task 
- Complete To Do task

(Data is in memory only and will be lost on switch off)

## API Endpoints

| **Method** | **URL**             | **Description**                   |
|------------|---------------------|-----------------------------------|
| GET        | `/api/todos`        | Retrieve all To Do items.          |
| GET        | `/api/todo/{id}`    | Retrieve a To Do item by ID.       |
| POST       | `/api/todo/`        | Create a new To Do item.           |
| PUT        | `/api/todo/{id}`    | Update a To Do item by ID.         |
| DELETE     | `/api/todo/{id}`    | Delete a To Do item by ID.         |

## HTML Interface

| **Method** | **URL**             | **Description**                   |
|------------|---------------------|-----------------------------------|
| GET        | `/todos`            | View all To Do items.              |
| GET        | `/todo/edit/{id}`   | Edit a To Do item by ID.           |


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