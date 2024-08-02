package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

type ToDoList struct {
	Items []ToDo `json:"todolist"`
}

type ToDo struct {
	Task string `json:"task"`
	Status string `json:"status"`
}

func NewToDo(name string) ToDo {
	return ToDo{
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

func (tdl *ToDoList) Update(task string, s string) (error) {
	index := slices.IndexFunc(tdl.Items, func(td ToDo) bool { return td.Task == task })
	if index < 0 {
		return errors.New("Task does not exist")
	}
	tdl.Items[index].Status = s
	return nil
}

func (tdl *ToDoList) Delete(task string) error {
	index := slices.IndexFunc(tdl.Items, func(td ToDo) bool { return td.Task == task })
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

func main() {
	// In memory To Do app
	reader := bufio.NewReader(os.Stdin)
	
	
	todos := ToDoList{}
	for {
		fmt.Println("Choose an operation: create, read, update, delete, or exit")
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		
		switch input {
		case "create":
			input = readCmd(*reader, "Enter a task: ")
			todos.Create(input)
			fmt.Println("task created")
		case "read":
			var list []ToDo
			list = todos.List()
			for _,i := range list {
				fmt.Printf("Task: %s; Status: %s\n", i.Task, i.Status)
			}
			
		case "update":
			task := readCmd(*reader, "Which task to update? ")
			status := readCmd(*reader, "Update task status to: ")
			err := todos.Update(task, status)
			if err != nil {
				println(err)
			} else {
				println("Task updated")
			}
		case "delete":
			task := readCmd(*reader, "Which task to delete? ")
			err := todos.Delete(task)
			if err != nil {
				println(err)
			} else {
				println("Task deleted")
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown operation")
		}
	}
}