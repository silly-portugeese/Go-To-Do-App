package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-webapp/backend/models"
	"todo-webapp/backend/service"
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

// --- helpers ---

func jsonResponse(writer http.ResponseWriter, data interface{}, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func getPathId(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	return strconv.Atoi(idStr)
}

// --- end helpers ---

func (h APIHandlers) FindAllHandler(writer http.ResponseWriter, request *http.Request) {
	items := h.Service.ListAllTItems()
	tdl := ToDoList{Count: len(items), Items: items}
	jsonResponse(writer, tdl, http.StatusOK)
}

func (h APIHandlers) FindByIdHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		jsonResponse(writer, map[string]string{"error": "Invalid ID: ID should be a number"}, http.StatusBadRequest)
		return
	}

	item, err := h.Service.GetItemById(id)
	if err != nil {
		// Assume it's a 404 for now
		jsonResponse(writer, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	jsonResponse(writer, item, http.StatusOK)

}

func (h APIHandlers) CreateHandler(writer http.ResponseWriter, request *http.Request) {

	var options map[string]string

	if err := json.NewDecoder(request.Body).Decode(&options); err != nil {
		jsonResponse(writer, map[string]string{"error": "failed to decode request"}, http.StatusBadRequest)
		return
	}

	task, ok := options["task"]
	if !ok {
		jsonResponse(writer, map[string]string{"error": "missing task"}, http.StatusBadRequest)
		return
	}

	item, err := h.Service.CreateItem(task)
	if err != nil {
		jsonResponse(writer, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	jsonResponse(writer, item, http.StatusOK)

}

func (h APIHandlers) UpdateHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		jsonResponse(writer, map[string]string{"error": "invalid ID: ID should be a number"}, http.StatusBadRequest)
		return
	}

	var options updateParams
	if err := json.NewDecoder(request.Body).Decode(&options); err != nil {
		jsonResponse(writer, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if (options.Task == "") && (options.Status == "") {
		jsonResponse(writer, map[string]string{"error": "at least one field (task or status) must be provided"}, http.StatusBadRequest)
		return
	}

	item, err := h.Service.UpdateItem(id, options.Task, options.Status)
	if err != nil {
		// Assume it's a 404 for now
		jsonResponse(writer, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	jsonResponse(writer, item, http.StatusOK)

}

func (h APIHandlers) DeleteHandler(writer http.ResponseWriter, request *http.Request) {

	id, err := getPathId(request)

	if err != nil {
		jsonResponse(writer, map[string]string{"error": "invalid ID: ID should be a number"}, http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteItem(id)
	if err != nil {
		// Assume it's a 404 for now
		jsonResponse(writer, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
