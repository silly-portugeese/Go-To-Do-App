package main

import (
	"fmt"
	"net/http"
	"todo-webapp/frontend/handlers"
)

func main() {

	// Initialize HTML handlers
	htmlHandlers := &handlers.HTMLHandlers{
		Host: "http://localhost:8001",
	}

	mux := http.NewServeMux()

	// HTML endpoints

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/*", http.StripPrefix("/static/", fs))

	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)

	mux.HandleFunc("GET /todos", htmlHandlers.InteractHandler)
	mux.HandleFunc("POST /todo/create", htmlHandlers.CreateHandler)
	mux.HandleFunc("GET /todo/edit/{id}", htmlHandlers.EditHandler)
	mux.HandleFunc("POST /todo/update/{id}", htmlHandlers.UpdateHandler)
	mux.HandleFunc("DELETE /todo/delete/{id}", htmlHandlers.DeleteHandler)

	port := "8002"

	fmt.Printf("Frontend server is listening on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}

}
