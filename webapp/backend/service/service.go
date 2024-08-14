package service

import (
	"errors"
	"strings"
	"todo-webapp/backend/models"
	"todo-webapp/backend/storage"
)

type Service struct {
	Store storage.IToDoRepository
}

func NewService(store storage.IToDoRepository) Service {
	return Service{store}
}

func (s *Service) StringToStatus(str string) (models.Status, error) {

	statusMapping := map[string]models.Status{
		"pending":     models.PENDING,
		"in progress": models.IN_PROGRESS,
		"completed":   models.COMPLETED,
	}

	str = strings.TrimSpace(str)
	str = strings.ToLower(str)

	status, ok := statusMapping[str]
	if ok {
		return status, nil
	} else {
		return status, errors.New("invalid status")
	}
}

func (s *Service) ListAllTItems() []models.ToDo {
	return s.Store.FindAll()
}

func (s *Service) GetItemById(id int) (models.ToDo, error) {
	return s.Store.FindById(id)
}

func (s *Service) CreateItem(task string) (models.ToDo, error) {

	task = strings.TrimSpace(task)
	if task == "" {
		return models.ToDo{}, errors.New("task is empty")
	}
	// TODO: pass the user id
	params := models.TodoCreateParams{Task: task, Status: models.PENDING, UserId: 1}
	item := s.Store.Create(params)
	return item, nil
}

func (s *Service) UpdateItem(id int, task string, statusStr string) (models.ToDo, error) {

	task = strings.TrimSpace(task)
	statusStr = strings.TrimSpace(statusStr)

	// Check that at least one field is provided and neither is an empty string
	if (task == "") && (statusStr == "") {
		return models.ToDo{}, errors.New("at least one field (task or status) must be provided")
	}

	var params models.TodoUpdateParams

	// Set task if not empty
	if task != "" {
		params.Task = &task
	}

	// Set status if not empty and valid
	if statusStr != "" {

		status, err := s.StringToStatus(statusStr)
		if err != nil {
			return models.ToDo{}, err
		}

		params.Status = &status
	}

	return s.Store.Update(id, params)
}

func (s *Service) DeleteItem(id int) error {
	return s.Store.Delete(id)
}
