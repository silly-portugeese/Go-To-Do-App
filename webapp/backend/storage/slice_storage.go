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

// Implementation
type todoStore struct {
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

// NewInMemoryStore creates and initializes a new in memory store for TODO items, returning a new `todoStore` instance.
// It launches a goroutine to manage commands and process requests asynchronously.
func NewInMemoryStore() todoStore {
	store := todoStore{items: []models.ToDo{}, nextId: 1, cmds: make(chan command)}
	go store.launchRequestManager()
	return store
}

// NewPrePopulatedInMemoryStore is similar to `NewEmptyInMemoryStore`, but sets up a new `todoStore` instance pre-populated with initial TODO items
func NewPrePopulatedInMemoryStore() todoStore {

	list := []models.ToDo{
		{Id: 1, Task: "Buy groceries", Status: models.PENDING, UserId: 1},
		{Id: 2, Task: "Write blog post", Status: models.IN_PROGRESS, UserId: 1},
		{Id: 3, Task: "Clean the house", Status: models.COMPLETED, UserId: 2},
		{Id: 4, Task: "Pay bills", Status: models.PENDING, UserId: 2},
		{Id: 5, Task: "Read a book", Status: models.COMPLETED, UserId: 3},
		{Id: 6, Task: "Prepare presentation", Status: models.IN_PROGRESS, UserId: 3},
		{Id: 7, Task: "Exercise", Status: models.PENDING, UserId: 3},
		{Id: 8, Task: "Call parents", Status: models.COMPLETED, UserId: 3},
		{Id: 9, Task: "Plan vacation", Status: models.PENDING, UserId: 4},
		{Id: 10, Task: "Learn Go", Status: models.IN_PROGRESS, UserId: 4},
	}

	store := todoStore{items: list, nextId: len(list) + 1, cmds: make(chan command)}
	go store.launchRequestManager()
	return store

}

func (tds *todoStore) launchRequestManager() {

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

func (tds *todoStore) executeFindAll(cmd command) {
	cmd.reply <- response{items: tds.items}
}

func (tds *todoStore) executeFindById(cmd command) {

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

func (tds *todoStore) executeCreate(cmd command) {

	params, hasParams := cmd.args["params"].(models.TodoCreateParams)
	
	if !hasParams {
		cmd.reply <- response{item: models.ToDo{}, err: errors.New("missing arguments")}
		return
	}

	item := models.ToDo{Id: tds.nextId, Task: params.Task, Status: params.Status, UserId: params.UserId}
	tds.items = append(tds.items, item)
	tds.nextId++
	cmd.reply <- response{item: item}
}

func (tds *todoStore) executeUpdate(cmd command) {
	
	id, hasId := cmd.args["id"].(int)
	params, hasParams := cmd.args["params"].(models.TodoUpdateParams)

	if !hasId {
		cmd.reply <- response{item: models.ToDo{}, err: fmt.Errorf("missing id: %d", id)}
		return
	}

	if !hasParams {
		cmd.reply <- response{item: models.ToDo{}, err: errors.New("missing arguments")}
		return
	}

	index := tds.getItemIndex(id)
	if index < 0 {
		cmd.reply <- response{item: models.ToDo{}, err: errors.New("to do item not found")}
		return
	}

	// simulate optional parameters by using pointers.
	if params.Task != nil {
		tds.items[index].Task = *params.Task
	}

	if params.Status != nil {
		tds.items[index].Status = *params.Status
	}
	
	cmd.reply <- response{item: tds.items[index], err: nil}
}

func (tds *todoStore) executeDelete(cmd command) {

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

func (tds *todoStore) getItemIndex(id int) int {
	return slices.IndexFunc(tds.items, func(td models.ToDo) bool { return td.Id == id })
}

// --- Interface Methods ---
// They are exposed to the rest of the application.
// Each one sends a task to the RequestManager so the request can be processed

func (tds todoStore) FindAll() []models.ToDo {
	ch := make(chan response)
	tds.cmds <- command{action: "list", reply: ch}
	response := <-ch
	return response.items
}

func (tds todoStore) FindById(id int) (models.ToDo, error) {
	ch := make(chan response)
	tds.cmds <- command{action: "read", reply: ch, args: map[string]interface{}{"id": id}}
	response := <-ch
	return response.item, response.err
}

func (tds todoStore) Create(params models.TodoCreateParams) models.ToDo {
	ch := make(chan response)
	tds.cmds <- command{action: "create", reply: ch, args: map[string]interface{}{"params":  params}}
	response := <-ch
	return response.item
}

func (tds todoStore) Update(id int, params models.TodoUpdateParams) (models.ToDo, error) {
	ch := make(chan response)
	tds.cmds <- command{action: "update", reply: ch, args: map[string]interface{}{"id": id, "params": params}}
	response := <-ch
	return response.item, response.err

}

func (tds todoStore) Delete(id int) error {
	ch := make(chan response)
	tds.cmds <- command{action: "delete", reply: ch, args: map[string]interface{}{"id": id}}
	response := <-ch
	return response.err
}
