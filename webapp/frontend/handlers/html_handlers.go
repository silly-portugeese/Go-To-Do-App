package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ToDo struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type ToDoList struct {
	Count int    `json:"count"`
	Items []ToDo `json:"todolist"`
}

type HTMLHandlers struct {
	Host string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal()
	}
}

func (h *HTMLHandlers) InteractHandler(writer http.ResponseWriter, request *http.Request) {
	requestURL := fmt.Sprintf("%s/api/todos", h.Host)

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatalln(err)
		// fmt.Printf("client: could not create request: %s\n", err)
		// os.Exit(1)
	}

	defer resp.Body.Close()

	// Check if the response status is OK.
	if resp.StatusCode != http.StatusOK {
		http.Error(writer, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	// Parse the JSON response.
	var data ToDoList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(writer, "Failed to parse data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/view.html")
	errorCheck(err)

	err = tmpl.Execute(writer, data)
	// errorCheck(err)
	if err != nil {
		// Handle the error if any and return
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *HTMLHandlers) DeleteHandler(writer http.ResponseWriter, request *http.Request) {

	id := request.PathValue("id")
	requestURL := fmt.Sprintf("%s/api/todo/delete/%s", h.Host, id)

	// Create the DELETE request
	req, err := http.NewRequest("DELETE", requestURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Optionally, set headers if required by the API
	// req.Header.Set("Accept", "application/json")

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
}

// func newHandler(writer http.ResponseWriter, request *http.Request) {
// 	tmpl, err := template.ParseFiles("new.html")
// 	errorCheck(err)
// 	err = tmpl.Execute(writer, nil)
// }

// func createHandler(writer http.ResponseWriter, request *http.Request) {
// 	todo := request.FormValue("todo")
// 	td := NewToDo(todo)
// 	tdl = append(tdl, td)

// 	http.Redirect(writer, request, "/interact", http.StatusFound)
// }
