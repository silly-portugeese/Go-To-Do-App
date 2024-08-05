package handlers

import (
	"encoding/json"

	"net/http"
	"strconv"
	"todo-webapp/models"
	"todo-webapp/service"
)

type ToDoList struct {
	Count int           `json:"count"`
	Items []models.ToDo `json:"todolist"`
}

type APIHandlers struct {
	Service *service.Service
}

type updateParams struct {
	Task   string
	Status string
}

// func write(writer http.ResponseWriter, msg string) {
// 	_, err := writer.Write([]byte(msg))
// 	if err != nil {
// 		log.Fatal()
// 	}
// }

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getPathId(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	return strconv.Atoi(idStr)
}

func (h APIHandlers) GetAllToDosHandler(writer http.ResponseWriter, request *http.Request) {
	items := h.Service.ListAllTItems()
	tdl := ToDoList{Count: len(items), Items: items}
	jsonResponse(writer, tdl)
}

func (h APIHandlers) GetToDoByIdHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		// http.Error(writer, "ID must be a number", http.StatusBadRequest)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Find item
	item, err := h.Service.GetItemById(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse(writer, item)

}

func (h APIHandlers) AddToDoHandler(writer http.ResponseWriter, request *http.Request) {

	var options map[string]string

	if err := json.NewDecoder(request.Body).Decode(&options); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	task, ok := options["task"]
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := h.Service.CreateItem(task)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	jsonResponse(writer, item)

}

func (h APIHandlers) UpdateToDoHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		// http.Error(writer, "ID must be a number", http.StatusBadRequest)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var options updateParams
	if err := json.NewDecoder(request.Body).Decode(&options); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		// http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if (options.Task == "") && (options.Status == "") {
		// 	return models.ToDo{}, errors.New("at least one field (task or status) must be provided")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := h.Service.UpdateItem(id, options.Task, options.Status)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonResponse(writer, item)

	// if status != "" && task != ""  {

	// }

	// var status *models.Status
	// if options.Status != "" {
	// 	status, err = h.Service.StringToStatus(options.Status)

	// 	if err != nil {
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// }

	// var item models.ToDo
	// if status != "" &&  options.Task != ""  {
	// 	item, err = h.Service.UpdateItem(id, &options.Task, status)
	// 	if err != nil {
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 	}
	// 	jsonResponse(writer, item)
	// } else if status != "" {
	// 	item, err := h.Service.UpdateItem(id, nil, status)
	// 	if err != nil {
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 	}
	// 	jsonResponse(writer, item)
	// } else {
	// 	item, err := h.Service.UpdateItem(id, &options.Task, nil)
	// 	if err != nil {
	// 		writer.WriteHeader(http.StatusBadRequest)
	// 	}
	// 	jsonResponse(writer, item)
	// }

}

func (h APIHandlers) DeleteToDoHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		// http.Error(writer, "ID must be a number", http.StatusBadRequest)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteItem(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
