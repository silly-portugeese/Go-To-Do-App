package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ToDos struct {
	List []string `json:"todolist"`
}

func print(l ...string) {
	for _, x := range l {
		fmt.Println(x)
	}
}

func (td ToDos) toJson() []byte {
	data, err := json.Marshal(td)
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

func readToDoJson(fname string) ToDos {
	var result ToDos

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

	tasks := ToDos{
		List: []string{
			"Buy groceries",
			"Write blog post",
			"Clean the house",
			"Pay bills",
			"Read a book",
			"Prepare presentation",
			"Exercise",
			"Call parents",
			"Plan vacation",
			"Learn Go",
		}}

	print(tasks.List...)

	data := tasks.toJson()
	fmt.Println(string(data))

	writeToFile("testj.json", tasks.toJson())

	result := readToDoJson("todos.json")
	fmt.Printf("%#v\n", result)

}

