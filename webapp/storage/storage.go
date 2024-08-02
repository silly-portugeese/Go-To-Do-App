package storage

import (
	"errors"
	"slices"
	"todo-webapp/models"
)

type IToDoRepository interface {
	FindAll() ([]models.ToDo)
	FindById(int) (models.ToDo, error)
	Create(string, models.Status) models.ToDo
	Update(int, string, models.Status) (models.ToDo, error)
	Delete(int) error
}

// Implementation
type ToDoStoreImpl struct {
	items  []models.ToDo
	nextId int
}

func NewInMemoryStore() ToDoStoreImpl {
	return ToDoStoreImpl{items: []models.ToDo{}, nextId: 1}
}

func (tds *ToDoStoreImpl) FindAll() []models.ToDo {
	return tds.items
}

func (tds *ToDoStoreImpl) FindById(id int) (models.ToDo, error) {
	index := tds.getItemIndex(id)
	if index < 0 {
		return models.ToDo{}, errors.New("to do item not found")
	}
	return tds.items[index], nil
}

func (tds *ToDoStoreImpl) Create(task string, status models.Status) models.ToDo {
	item := models.ToDo{Id: tds.nextId, Task: task, Status: status}
	tds.items = append(tds.items, item)
	tds.nextId += 1
	return item
}

func (tds *ToDoStoreImpl) Update(id int, task string, status models.Status) (models.ToDo, error) {

	index := tds.getItemIndex(id)

	if index < 0 {
		return models.ToDo{}, errors.New("to do item not found")
	}

	tds.items[index].Task = task
	tds.items[index].Status = status

	return tds.items[index], nil

}

func (tds *ToDoStoreImpl) Delete(id int) error {
	index := tds.getItemIndex(id)

	if index < 0 {
		return errors.New("to do item not found")
	}

	tds.items = append(tds.items[:index], tds.items[index+1:]...)
	return nil
}

func (tds *ToDoStoreImpl) getItemIndex(id int) int {
	return slices.IndexFunc(tds.items, func(td models.ToDo) bool { return td.Id == id })
}
