package models

type Status string

const (
	PENDING     Status = "Pending"
	IN_PROGRESS Status = "In Progress"
	COMPLETED   Status = "Completed"
)

type ToDo struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status Status `json:"status"`
	UserId int    `json:"userid"`
}

type ToDoList struct {
	Count int    `json:"count"`
	Items []ToDo `json:"todolist"`
}

// TodoCreateData encapsulates the parameters needed to create a new todo.
type TodoCreateData struct {
	Task   string
	Status Status
	UserId int
}

// TodoUpdateData encapsulates the parameters needed to update an existing todo.
type TodoUpdateData struct {
	Task   *string // Use pointers to allow partial updates
	Status *Status
}
