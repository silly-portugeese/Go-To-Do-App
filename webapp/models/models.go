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
}