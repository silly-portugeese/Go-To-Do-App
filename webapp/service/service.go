package main

import (
	"errors"
	"strings"
	"todo-webapp/models"
	"todo-webapp/storage"
)

type Service struct {
    store storage.IToDoRepository
}

func NewService(store storage.IToDoRepository) *Service {
    return &Service{store}
}


func (s *Service) ListAllTItems() ([]models.ToDo) {
	return s.store.FindAll()
}

func (s *Service) GetItemById(id int ) (models.ToDo, error) {
	return s.store.FindById(id)
}

func (s *Service) CreateItem(task string) (models.ToDo, error) {
	
	task = strings.TrimSpace(task)
	if task == "" {
		return models.ToDo{}, errors.New("task is empty")
	}
    item := s.store.Create(task, models.PENDING)
    return item, nil
}

func (s *Service) UpdateItem(id int, task string, status models.Status) (models.ToDo, error) {
	return s.store.Update(id, task, status)
}

func (s *Service) DeleteItem(id int) (error) {
	return s.store.Delete(id)
}