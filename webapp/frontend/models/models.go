package models

type ToDo struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type ToDoList struct {
	Count int    `json:"count"`
	Items []ToDo `json:"todolist"`
}

