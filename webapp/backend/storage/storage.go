package storage

import (
	"errors"
	"fmt"
	"slices"
	"todo-webapp/backend/models"
)

// Understanding Channel Types
// Send-Only Channel (chan<- T): This type of channel can only be used to send data.
// Receive-Only Channel (<-chan T): This type of channel can only be used to receive data.
// Bidirectional Channel (chan T): This type of channel can be used for both sending and receiving data.

type IToDoRepository interface {
	FindAll() []models.ToDo
	FindById(int) (models.ToDo, error)
	Create(string, models.Status) models.ToDo
	Update(int, *string, *models.Status) (models.ToDo, error)
	Delete(int) error
}

// Implementation
type toDoStoreImpl struct {
	cmds   chan command
	items  []models.ToDo
	nextId int
}

type response struct {
	item  models.ToDo
	items []models.ToDo
	err   error
}

type command struct {
	action string
	args   map[string]interface{}
	reply  chan response
}

func NewInMemoryStore() toDoStoreImpl {

	list := []models.ToDo{
		{Id: 1, Task: "Buy groceries", Status: "Pending"},
		{Id: 2, Task: "Write blog post", Status: "In Progress"},
		{Id: 3, Task: "Clean the house", Status: "Completed"},
		{Id: 4, Task: "Pay bills", Status: "Pending"},
		{Id: 5, Task: "Read a book", Status: "Completed"},
		{Id: 6, Task: "Prepare presentation", Status: "In Progress"},
		{Id: 7, Task: "Exercise", Status: "Pending"},
		{Id: 8, Task: "Call parents", Status: "Completed"},
		{Id: 9, Task: "Plan vacation", Status: "Pending"},
		{Id: 10, Task: "Learn Go", Status: "In Progress"},
	}

	store := toDoStoreImpl{items: list, nextId: len(list) + 1, cmds: make(chan command)}
	go store.launchRequestManager()
	return store

}

func NewEmptyInMemoryStore() toDoStoreImpl {
	store := toDoStoreImpl{items: []models.ToDo{}, nextId: 1, cmds: make(chan command)}
	go store.launchRequestManager()
	return store
}

func (tds *toDoStoreImpl) launchRequestManager() {

	executers := map[string]func(cmd command){
		"create": tds.executeCreate,
		"read":   tds.executeFindById,
		"update": tds.executeUpdate,
		"delete": tds.executeDelete,
		"list":   tds.executeFindAll,
	}

	for cmd := range tds.cmds {
		execute, ok := executers[cmd.action]
		if !ok {
			cmd.reply <- response{err: fmt.Errorf("unknown action: %s", cmd.action)}
			continue
		}
		execute(cmd)
	}
}

// --- Internal Methods ---
// These methods are used internally by the RequestManager to handle specific commands: create, read, update, delete, list

func (tds *toDoStoreImpl) executeFindAll(cmd command) {
	cmd.reply <- response{items: tds.items}
}

func (tds *toDoStoreImpl) executeFindById(cmd command) {

	id, hasId := cmd.args["id"].(int)

	if !hasId {
		cmd.reply <- response{item: models.ToDo{}, err: fmt.Errorf("missing id: %d", id)}
		return
	}

	index := tds.getItemIndex(id)
	if index < 0 {
		cmd.reply <- response{item: models.ToDo{}, err: errors.New("to do item not found")}
		return
	}

	cmd.reply <- response{item: tds.items[index], err: nil}
}

func (tds *toDoStoreImpl) executeCreate(cmd command) {

	task, hasTask := cmd.args["task"].(string)
	status, hasStatus := cmd.args["status"].(models.Status)

	if !hasTask || !hasStatus {
		cmd.reply <- response{err: errors.New("missing arguments")}
		return
	}

	item := models.ToDo{Id: tds.nextId, Task: task, Status: status}
	tds.items = append(tds.items, item)
	tds.nextId++
	cmd.reply <- response{item: item}
}

func (tds *toDoStoreImpl) executeUpdate(cmd command) {

	id, hasId := cmd.args["id"].(int)
	task, _ := cmd.args["task"].(*string)
	status, _ := cmd.args["status"].(*models.Status)

	if !hasId {
		cmd.reply <- response{item: models.ToDo{}, err: fmt.Errorf("missing id: %d", id)}
		return
	}

	index := tds.getItemIndex(id)
	if index < 0 {
		cmd.reply <- response{item: models.ToDo{}, err: errors.New("to do item not found")}
		return
	}

	// simulate optional parameters by using pointers.
	if task != nil {
		tds.items[index].Task = *task
	}

	if status != nil {
		tds.items[index].Status = *status
	}

	cmd.reply <- response{item: tds.items[index], err: nil}
}

func (tds *toDoStoreImpl) executeDelete(cmd command) {

	id, hasId := cmd.args["id"].(int)
	if !hasId {
		cmd.reply <- response{err: fmt.Errorf("missing id: %d", id)}
		return
	}

	index := tds.getItemIndex(id)
	if index < 0 {
		cmd.reply <- response{err: errors.New("to do item not found")}
		return
	}

	tds.items = append(tds.items[:index], tds.items[index+1:]...)
	cmd.reply <- response{err: nil}
}

func (tds *toDoStoreImpl) getItemIndex(id int) int {
	return slices.IndexFunc(tds.items, func(td models.ToDo) bool { return td.Id == id })
}

// --- Interface Methods ---
// They are exposed to the rest of the application.

func (tds *toDoStoreImpl) FindAll() []models.ToDo {
	ch := make(chan response)
	tds.cmds <- command{action: "list", reply: ch}
	response := <-ch
	return response.items
}

func (tds *toDoStoreImpl) FindById(id int) (models.ToDo, error) {
	ch := make(chan response)
	tds.cmds <- command{action: "read", reply: ch, args: map[string]interface{}{"id": id}}
	response := <-ch
	return response.item, response.err
}

func (tds *toDoStoreImpl) Create(task string, status models.Status) models.ToDo {
	ch := make(chan response)
	tds.cmds <- command{action: "create", reply: ch, args: map[string]interface{}{"task": task, "status": status}}
	response := <-ch
	return response.item
}

func (tds *toDoStoreImpl) Update(id int, task *string, status *models.Status) (models.ToDo, error) {
	ch := make(chan response)
	tds.cmds <- command{action: "update", reply: ch, args: map[string]interface{}{"id": id, "task": task, "status": status}}
	response := <-ch
	return response.item, response.err

}

func (tds *toDoStoreImpl) Delete(id int) error {
	ch := make(chan response)
	tds.cmds <- command{action: "delete", reply: ch, args: map[string]interface{}{"id": id}}
	response := <-ch
	return response.err
}
