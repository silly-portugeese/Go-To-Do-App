package storage

import (
	"reflect"
	"testing"
	"todo-webapp/backend/models"
)

// File generated using option "Go: generate unit tests for function"

// if your working directory is that of your package, to test all of the files you could run:
// go test ./...

// if you wanted to get test coverage, you could run:
// go test ./... -cover

var todoSlice []models.ToDo

func Init() {
	todoSlice = []models.ToDo{
		{Id: 1, Task: "Buy groceries", Status: models.PENDING},
		{Id: 2, Task: "Write blog post", Status: models.IN_PROGRESS},
		{Id: 3, Task: "Clean the house", Status: models.COMPLETED},
		{Id: 4, Task: "Pay bills", Status: models.PENDING},
		{Id: 5, Task: "Read a book", Status: models.COMPLETED},
		{Id: 6, Task: "Prepare presentation", Status: models.IN_PROGRESS},
		{Id: 7, Task: "Exercise", Status: models.PENDING},
		{Id: 8, Task: "Call parents", Status: models.COMPLETED},
		{Id: 9, Task: "Plan vacation", Status: models.PENDING},
		{Id: 10, Task: "Learn Go", Status: models.IN_PROGRESS},
	}
}

func Test_todoStore_FindAll(t *testing.T) {

	Init()
	type fields struct {
		cmds   chan command
		items  []models.ToDo
		nextId int
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.ToDo
	}{
		{
			name:   "Valid to do list",
			fields: fields{items: todoSlice, nextId: len(todoSlice) + 1},
			want:   todoSlice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := &todoStore{
				cmds:   make(chan command),
				items:  tt.fields.items,
				nextId: tt.fields.nextId,
			}
			go tds.launchRequestManager()
			if got := tds.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoStore.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoStore_FindById(t *testing.T) {

	Init()
	type fields struct {
		cmds   chan command
		items  []models.ToDo
		nextId int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ToDo
		wantErr bool
	}{
		{
			name:    "Return existing To do item",
			fields:  fields{items: todoSlice, nextId: len(todoSlice) + 1},
			args:    args{id: 3},
			want:    models.ToDo{Id: 3, Task: "Clean the house", Status: "Completed"},
			wantErr: false,
		},
		{
			name:    "Item  not found",
			fields:  fields{items: todoSlice, nextId: len(todoSlice) + 1},
			args:    args{id: 456},
			want:    models.ToDo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := &todoStore{
				cmds:   make(chan command),
				items:  tt.fields.items,
				nextId: tt.fields.nextId,
			}
			go tds.launchRequestManager()
			got, err := tds.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoStore.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoStore.FindById() = %v, want %v", got, tt.want)
			}
			close(tds.cmds)
		})
	}
}

func Test_todoStore_Create(t *testing.T) {

	Init()
	type fields struct {
		cmds   chan command
		items  []models.ToDo
		nextId int
	}
	type args struct {
		task   string
		status models.Status
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.ToDo
	}{
		{
			name:   "Create a valid To Do",
			fields: fields{items: []models.ToDo{}, nextId: 1},
			args:   args{task: "Sing", status: models.IN_PROGRESS},
			want:   models.ToDo{Id: 1, Task: "Sing", Status: models.IN_PROGRESS},
		},
		{
			name:   "Valid next Id defined",
			fields: fields{items: []models.ToDo{}, nextId: 3},
			args:   args{task: "Sing", status: models.IN_PROGRESS},
			want:   models.ToDo{Id: 3, Task: "Sing", Status: models.IN_PROGRESS},
		},
	}

	// fmt.Println(todoSlice)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := &todoStore{
				cmds:   make(chan command),
				items:  tt.fields.items,
				nextId: tt.fields.nextId,
			}
			go tds.launchRequestManager()
			if got := tds.Create(tt.args.task, tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoStore.Create() = %v, want %v", got, tt.want)
			}
			close(tds.cmds)
		})
	}
}

func Test_todoStore_Update(t *testing.T) {

	Init()
	task := "Clean the floor"
	status := models.IN_PROGRESS

	type fields struct {
		cmds   chan command
		items  []models.ToDo
		nextId int
	}
	type args struct {
		id     int
		task   *string
		status *models.Status
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.ToDo
		wantErr bool
	}{
		{
			name:    "Unsuccessful update: item not found",
			fields:  fields{items: []models.ToDo{}, nextId: 1},
			args:    args{id: 3, task: &task, status: &status},
			want:    models.ToDo{},
			wantErr: true,
		},
		{
			name:    "Unsuccessful update: not enough paramenters",
			fields:  fields{items: todoSlice, nextId: len(todoSlice) + 1},
			args:    args{id: 4, task: nil, status: nil},
			want:    models.ToDo{Id: 4, Task: "Pay bills", Status: models.PENDING},
			wantErr: false,
		},
		{
			name:    "Successful update of item",
			fields:  fields{items: todoSlice, nextId: len(todoSlice) + 1},
			args:    args{id: 3, task: &task, status: &status},
			want:    models.ToDo{Id: 3, Task: "Clean the floor", Status: models.IN_PROGRESS},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := &todoStore{
				cmds:   make(chan command),
				items:  tt.fields.items,
				nextId: tt.fields.nextId,
			}
			go tds.launchRequestManager()
			got, err := tds.Update(tt.args.id, tt.args.task, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoStore.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoStore.Update() = %v, want %v", got, tt.want)
			}
			close(tds.cmds)
		})
	}
}

func Test_todoStore_Delete(t *testing.T) {

	Init()

	type fields struct {
		cmds   chan command
		items  []models.ToDo
		nextId int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Successful deletion",
			fields:  fields{items: todoSlice, nextId: len(todoSlice) + 1},
			args:    args{id: 3},
			wantErr: false,
		},
		{
			name:    "Unsuccessful deletion: item not found",
			fields:  fields{items: []models.ToDo{}, nextId: 1},
			args:    args{id: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tds := &todoStore{
				cmds:   make(chan command),
				items:  tt.fields.items,
				nextId: tt.fields.nextId,
			}
			go tds.launchRequestManager()
			if err := tds.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("todoStore.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
