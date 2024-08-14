package storage

import "todo-webapp/backend/models"

type IToDoRepository interface {
	FindAll() []models.ToDo
	FindById(int) (models.ToDo, error)
	Create(models.TodoCreateData) models.ToDo
	Update(int, models.TodoUpdateData) (models.ToDo, error)
	Delete(int) error
}
