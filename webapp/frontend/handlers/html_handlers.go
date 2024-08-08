package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"todo-webapp/frontend/models"
)

type HTMLHandlers struct {
	Host string
}


// --- helpers ---

func renderHTMLTemplate(writer http.ResponseWriter, data any, file string) {

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		respondWithInternalServerError(writer)
		return
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		respondWithInternalServerError(writer)
		return
	}
}

func renderHTMLText(writer http.ResponseWriter, data any, html string) {

	tmpl := template.Must(template.New("item").Parse(html))
	
	err := tmpl.Execute(writer, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		respondWithInternalServerError(writer)
		return
	}
}

func makeHttpRequest(method string, url string,  body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("Failed to create request: %v\n", err)
		return &http.Response{}, err		
	}

	// Set Content-Type header for JSON
	req.Header.Set("Accept", "application/json")

	// Send the request using http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to make request to %s: %v\n", url, err)
	}

	return resp, err

}

func respondWithInternalServerError(writer http.ResponseWriter) {
	http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
}

// --- end helpers ---



func (h *HTMLHandlers) InteractHandler(writer http.ResponseWriter, request *http.Request) {

	requestURL := fmt.Sprintf("%s/api/todos",h.Host)

	// Create the GET request
	resp, err := makeHttpRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		renderHTMLTemplate(writer, nil, "static/view.html")
		return
	}

	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch data from %s. HTTP status code: %d\n", requestURL, resp.StatusCode)
		renderHTMLTemplate(writer, nil, "static/view.html")
		return
	}

	// Parse the JSON response.
	var apiData models.ToDoList
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		log.Printf("Failed to parse data: %v\n", err)
		respondWithInternalServerError(writer)
		return
	}

	renderHTMLTemplate(writer, apiData, "static/view.html")
}



func (h *HTMLHandlers) EditHandler(writer http.ResponseWriter, request *http.Request) {

	id := request.PathValue("id")
	requestURL := fmt.Sprintf("%s/api/todo/%s", h.Host, id)
	
	// Create the GET request
	resp, err := makeHttpRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		respondWithInternalServerError(writer)
		return
	}

	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch data from %s. HTTP status code: %d\n", requestURL, resp.StatusCode)
		http.Error(writer, resp.Status, resp.StatusCode)
		return
	}

	// Parse the JSON response.
	var apiData models.ToDo
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		log.Printf("Failed to parse data: %v\n", err)
		respondWithInternalServerError(writer)
		return
	}

	renderHTMLTemplate(writer, apiData, "static/edit.html")
}



func (h *HTMLHandlers) UpdateHandler(writer http.ResponseWriter, request *http.Request) {

	id := request.PathValue("id")
	requestURL := fmt.Sprintf("%s/api/todo/update/%s", h.Host, id)

	task := request.FormValue("task")
	status := request.FormValue("status")
	data := map[string]string{"task": task, "status": status}

	// Marshal the payload into JSON
	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to parse data: %v\n", err)
		respondWithInternalServerError(writer)
		return
	}

	// Create the PUT request
	resp, err := makeHttpRequest(http.MethodPut, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		respondWithInternalServerError(writer)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to update item. HTTP status code: %d", resp.StatusCode)
		http.Error(writer, resp.Status, resp.StatusCode)
		return
	}

	http.Redirect(writer, request, "/todos", http.StatusFound)
}



func (h *HTMLHandlers) CreateHandler(writer http.ResponseWriter, request *http.Request) {

	requestURL := fmt.Sprintf("%s/api/todo/create", h.Host)

	task := request.FormValue("task")
	data := map[string]string{"task": task}

	// Marshal the payload into JSON
	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to parse data: %v\n", err)
		respondWithInternalServerError(writer)
		return
	}

	// Create the POST request
	resp, err := makeHttpRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		respondWithInternalServerError(writer)
		return
	}

	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to create item. HTTP status code: %d", resp.StatusCode)
		http.Error(writer, resp.Status, resp.StatusCode)
		return
	}

	// Parse the JSON response.
	var apiData models.ToDo
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		log.Printf("Failed to parse data: %v\n", err)
		respondWithInternalServerError(writer)
		return
	}

	// Render the new item HTML to be inserted into the list
	renderHTMLText(
		writer,
		apiData,
		`
		<li id="todo-item-{{ .Id }}" class="tod0-item">
		<div class="task">{{.Task}}</div>
		<div class="status" style="background-color: #f0ac00;">{{.Status}}</div>
		<div>
			<a href="/todo/edit/{{ .Id }}">✏️</a>
			<button hx-delete="/todo/delete/{{ .Id }}" hx-target="#todo-item-{{ .Id }}" hx-swap="outerHTML">❌</button>
		</div>  
	</li>
	`)
}



func (h *HTMLHandlers) DeleteHandler(writer http.ResponseWriter, request *http.Request) {

	id := request.PathValue("id")
	requestURL := fmt.Sprintf("%s/api/todo/delete/%s", h.Host, id)

	// Create the DELETE request
	resp, err := makeHttpRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		respondWithInternalServerError(writer)
		return
	}

	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusNoContent {
		log.Printf("Failed to delete item. HTTP status code: %d", resp.StatusCode)
		http.Error(writer, resp.Status, resp.StatusCode)
		return
	}
}
