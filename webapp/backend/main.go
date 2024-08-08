package main

import (
	"fmt"
	"net/http"
	"todo-webapp/backend/handlers"
	"todo-webapp/backend/service"
	"todo-webapp/backend/storage"
)

func main() {

	store := storage.NewInMemoryStore()
	// store := storage.NewEmptyInMemoryStore()

	service := service.NewService(&store)

	// Initialize API handlers
	apiHandlers := handlers.APIHandlers{
		Service: service,
	}

	mux := http.NewServeMux()

	// JSON API endpoints
	mux.HandleFunc("GET /api/todos", apiHandlers.FindAllHandler)
	mux.HandleFunc("GET /api/todo/{id}", apiHandlers.FindByIdHandler)
	mux.HandleFunc("POST /api/todo/", apiHandlers.CreateHandler)
	mux.HandleFunc("PUT /api/todo/update/{id}", apiHandlers.UpdateHandler)
	mux.HandleFunc("DELETE /api/todo/delete/{id}", apiHandlers.DeleteHandler)

	port := "8001"

	fmt.Printf("Backend server is listening on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}
