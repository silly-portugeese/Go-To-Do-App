package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type ToDoList struct {
	List []ToDo `json:"todolist"`
}

type ToDo struct {
	Task string `json:"task"`
	Status string `json:"status"`
}

func print(l []string) {
	for _, x := range l {
		fmt.Println(x)
	}
}


func (td ToDoList) getTasks() []string {
	var list []string
    for _, item := range td.List {
        list = append(list, item.Task)
    }
    return list
}

func (td ToDoList) getStatuses() []string {
	var list []string
    for _, item := range td.List {
        list = append(list, item.Status)
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

	print(tasks.getTasks())

	data := toJson(tasks)
	fmt.Println(string(data))

	writeToFile("testj.json", toJson(tasks))

	result := readToDoJson("testj.json")
	fmt.Printf("%#v\n", result)

	// Print tasks and statuses concurrently
	taskCh := make(chan bool)
	statusCh := make(chan bool)
	
	var wg sync.WaitGroup

	wg.Add(2)
	
	go printTask(tasks.getTasks(), taskCh, statusCh, &wg)
	go printStatus(tasks.getStatuses(), taskCh, statusCh, &wg)

	wg.Wait()
	close(statusCh)
	close(taskCh)

}

func printTask(sl []string, ch1, ch2 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for _,s := range sl {
		fmt.Print("Task: ", s)
		ch2 <- true
		<-ch1 // wait for status to be printed
	}
}

func printStatus(sl []string, ch1, ch2 chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for _,s := range sl {
		<-ch2  // wait for the task to be printed
		fmt.Println("; Status:", s)
		ch1 <- true
	}
}

