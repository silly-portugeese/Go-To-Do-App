package storage

import (
	"strconv"
	"testing"
)

// https://medium.com/nerd-for-tech/benchmarking-your-solution-in-go-940b528416c
// go test -bench=.
// go test -bench=BenchmarkTest_todoStore_FindAll
// go test -bench=. -benchtime=20s

func BenchmarkTest_todoStore_FindAll(b *testing.B) {
    store := NewInMemoryStore()
    
    for i := 0; i < b.N; i++ {
        store.FindAll()
    }
}

func BenchmarkTest_todoStore_FindById(b *testing.B) {
    store := NewInMemoryStore()
    id := 10

    for i := 0; i < b.N; i++ {
        _, _ = store.FindById(id)
    }
}

func BenchmarkTest_todoStore_Create(b *testing.B) {
	store := NewInMemoryStore()

    for i := 0; i < b.N; i++ {
		task := "Task " + strconv.Itoa(i)
        store.Create(task,"Pending")
    }
}


func BenchmarkTest_todoStore_Update(b *testing.B) {
	store := NewInMemoryStore()
    id := 10
	task := "Do Something"

    for i := 0; i < b.N; i++ {
       _, _ = store.Update(id, &task, nil)
    }
}


func BenchmarkTest_todoStore_Delete(b *testing.B) {
	store := NewInMemoryStore()
	id := 9
    for i := 0; i < b.N; i++ {
        store.Delete(id)
    }
}