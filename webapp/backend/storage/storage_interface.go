package storage

import "todo-webapp/backend/models"

type IToDoRepository interface {
	FindAll() []models.ToDo
	FindById(int) (models.ToDo, error)
	Create(string, models.Status) models.ToDo
	Update(int, *string, *models.Status) (models.ToDo, error)
	Delete(int) error
}
