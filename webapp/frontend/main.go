package main

import (
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

	http.ListenAndServe(":8002", mux)
	
}
