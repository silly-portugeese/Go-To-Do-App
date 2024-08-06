package main

import (
	"net/http"
	"todo-webapp/backend/handlers"
	"todo-webapp/backend/service"
	"todo-webapp/backend/storage"
)

func main() {

	// store := storage.NewInMemoryStore()
	store := storage.NewEmptyInMemoryStore()

	service := service.NewService(&store)

	// Initialize API handlers
	apiHandlers := &handlers.APIHandlers{
		Service: service,
	}

	mux := http.NewServeMux()

	// JSON API endpoints
	mux.HandleFunc("GET /api/todos", apiHandlers.GetAllToDosHandler)
	mux.HandleFunc("GET /api/todo/{id}", apiHandlers.GetToDoByIdHandler)
	mux.HandleFunc("POST /api/todo/", apiHandlers.AddToDoHandler)
	mux.HandleFunc("PUT /api/todo/update/{id}", apiHandlers.UpdateToDoHandler)
	mux.HandleFunc("DELETE /api/todo/delete/{id}", apiHandlers.DeleteToDoHandler)

	http.ListenAndServe(":8001", mux)

	
}
