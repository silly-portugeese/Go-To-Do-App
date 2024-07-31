package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ToDoList struct {
	List []ToDo `json:"todolist"`
}

type ToDo struct {
	Task string `json:"task"`
	Status string `json:"status"`
}


func print(l ...string) {
	for _, x := range l {
		fmt.Println(x)
	}
}


func (td ToDoList) getList() []string {
	var list []string
    for _, item := range td.List {
        list = append(list, item.Task)
    }
    return list
}


func toJson(v any) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return data
}


func writeToFile(fname string, data []byte) error {

	f, err := os.Create(fname)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(data)

	if err != nil {
		return err
	}

	return nil

}

func readFile(fname string) ([]byte, error) {
	// f, err := os.Open(fname)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 1024)
	// data, err := f.Read(buf)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	data, err := os.ReadFile(fname)
	if err != nil {
		return []byte{}, err
	}

	return data, nil

}

func readToDoJson(fname string) ToDoList {
	var result ToDoList

	data, err := readFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

func main() {

	items := []ToDo{
		{"Buy groceries", "Pending"},
		{"Write blog post", "In Progress"},
		{"Clean the house", "Completed"},
		{"Pay bills", "Pending"},
		{"Read a book", "Completed"},
		{"Prepare presentation", "In Progress"},
		{"Exercise", "Pending"},
		{"Call parents", "Completed"},
		{"Plan vacation", "Pending"},
		{"Learn Go", "In Progress"},
	}

	
	tasks := ToDoList{List: items}

	print(tasks.getList()...)

	data := toJson(tasks)
	fmt.Println(string(data))

	writeToFile("testj.json", toJson(tasks))

	result := readToDoJson("testj.json")
	fmt.Printf("%#v\n", result)

}

