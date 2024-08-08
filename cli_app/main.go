package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var IdCounter int

type ToDoList struct {
	Items []ToDo `json:"todolist"`
}

type ToDo struct {
	Id int `json:"id"`
	Task string `json:"task"`
	Status string `json:"status"`
}

func NewToDo(name string) ToDo {
	IdCounter++

	return ToDo{
		Id: IdCounter,
		Task: name,
		Status:  "Pending",
	}
}


type IToDoInterface interface {
	Create(string) 
	List()
	Update()
	Delete()
}

func (tdl *ToDoList) Create(task string) ToDo {
	todo := NewToDo(task)
	tdl.Items = append(tdl.Items,  todo)
	return todo
}

func (tdl *ToDoList) List() ([]ToDo) {
	return tdl.Items
}

func (tdl *ToDoList) Read(id int) (ToDo, error) {
	index := slices.IndexFunc(tdl.Items, func(td ToDo) bool { return td.Id == id})
	if index < 0 {
		return ToDo{}, errors.New("Task does not exist")
	}
	return tdl.Items[index], nil
}

func (tdl *ToDoList) Update(id int, name string, status string) (error) {
	index := slices.IndexFunc(tdl.Items, func(td ToDo) bool { return td.Id == id})
	if index < 0 {
		return errors.New("Task does not exist")
	}

	if name == "" && status == "" {
		return errors.New("Nothing to update")
	}

	if name != "" {
		tdl.Items[index].Task = name
	}
	if status != "" {
		tdl.Items[index].Status = status
	}
	return nil
}

func (tdl *ToDoList) Delete(id int) error {
	index := slices.IndexFunc(tdl.Items, func(td ToDo) bool { return td.Id == id })
	if index < 0 {
		return errors.New("Task does not exist")
	}
	tdl.Items = append(tdl.Items[:index], tdl.Items[index+1:]...)
	return nil
}

func readCmd(reader bufio.Reader, msg string) (s string) {
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func strToInt(s string) {

}

func main() {
	// In memory To Do app
	reader := bufio.NewReader(os.Stdin)
	
	
	todos := ToDoList{}
	for {
		fmt.Println("Choose an operation: create (c), list (l), read (r), update (u), delete (d), or exit (e)")
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		
		switch input {
		case "create", "c":
			input = readCmd(*reader, "Enter a task: ")
			todos.Create(input)
			fmt.Println("Task created")
		
		case "list", "l":
			var list []ToDo
			list = todos.List()
			for _,i := range list {
				fmt.Printf("ID: %d; Task: %s; Status: %s\n", i.Id, i.Task, i.Status)
			}

		case "read", "r":
			idStr := readCmd(*reader, "Give task ID: ")

			id, err := strconv.Atoi(idStr); 
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			} 
			task, err := todos.Read(id)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("ID: %d; Task: %s; Status: %s\n", task.Id, task.Task, task.Status)
			}
			
			
		case "update", "u":
			idStr := readCmd(*reader, "Which task to update? Give task ID: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			
			_, err = todos.Read(id)
			if err != nil {
				fmt.Println(err.Error())
				continue
			} 

			task := readCmd(*reader, "Update task name to... If you don't want to update it, leave it empty: ")
			status := readCmd(*reader, "Update task status to... If you don't want to update it, leave it empty: ")

			err = todos.Update(id, task, status)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Task updated")
			}
		
		case "delete", "d":
			idStr := readCmd(*reader, "Which task to delete? Give task ID: ")
			id, err := strconv.Atoi(idStr); 
			
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}

			err = todos.Delete(id)
			if err != nil {
				fmt.Println(err.Error())
				
			} else {
				fmt.Println("Task deleted")
			}
		
		case "exit", "e":
			fmt.Println("Exiting...")
			return
		
		default:
			fmt.Println("Unknown operation")
		}
	}
}