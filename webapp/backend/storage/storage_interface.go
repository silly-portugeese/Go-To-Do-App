package storage

import "todo-webapp/backend/models"

type IToDoRepository interface {
	FindAll() []models.ToDo
	FindById(int) (models.ToDo, error)
	Create(models.TodoCreateParams) models.ToDo
	Update(int, models.TodoUpdateParams) (models.ToDo, error)
	Delete(int) error
}
