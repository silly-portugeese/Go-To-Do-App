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

// func write(writer http.ResponseWriter, msg string) {
// 	_, err := writer.Write([]byte(msg))
// 	errorCheck(err)
// }

// requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
// req, err := http.NewRequest(http.MethodGet, requestURL, nil)
// if err != nil {
// 	fmt.Printf("client: could not create request: %s\n", err)
// 	os.Exit(1)
// }

// res, err := http.DefaultClient.Do(req)
// if err != nil {
// 	fmt.Printf("client: error making http request: %s\n", err)
// 	os.Exit(1)
// }

// fmt.Printf("client: got response!\n")
// fmt.Printf("client: status code: %d\n", res.StatusCode)

// resBody, err := ioutil.ReadAll(res.Body)
// if err != nil {
// 	fmt.Printf("client: could not read response body: %s\n", err)
// 	os.Exit(1)
// }
// fmt.Printf("client: response body: %s\n", resBody)

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

// func main() {
// 	http.HandleFunc("/interact", interactHandler)
// 	http.HandleFunc("/new", newHandler)
// 	http.HandleFunc("/create", createHandler)

// 	// http.HandleFunc("/update/{id}", updateHandler)
// 	// http.HandleFunc("/delete/{id}", deleteHandler)
// 	// http.HandleFunc("/create", createHandler)

// 	err := http.ListenAndServe("localhost:8080", nil)
// 	log.Fatal(err)

// }
