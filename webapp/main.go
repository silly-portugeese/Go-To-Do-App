package main

import (
	"net/http"
	"todo-webapp/handlers"
	"todo-webapp/storage"
	"todo-webapp/service"
)


func main() {

	store := storage.NewInMemoryStore()
    service := service.NewService(&store)


    // Initialize API handlers
    apiHandlers := &handlers.APIHandlers{
        Service: service,
    }

	// JSON API endpoints
	http.HandleFunc("GET /api/todos", apiHandlers.GetAllToDosHandler)
	http.HandleFunc("GET /api/todo/{id}", apiHandlers.GetToDoByIdHandler)
	http.HandleFunc("POST /api/todo/", apiHandlers.AddToDoHandler)
	http.HandleFunc("PUT /api/todo/update/{id}", apiHandlers.UpdateToDoHandler)
	http.HandleFunc("DELETE /api/todo/delete/{id}", apiHandlers.DeleteToDoHandler)

	http.ListenAndServe(":8080", nil)
}