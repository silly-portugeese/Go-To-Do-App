package storage

import (
 	"strconv"
	"testing"
)

// https://medium.com/nerd-for-tech/benchmarking-your-solution-in-go-940b528416c
// go test -bench=.
// go test -bench=BenchmarkTestToDoStoreImpl_FindAll
// go test -bench=. -benchtime=20s

func BenchmarkTestToDoStoreImpl_FindAll(b *testing.B) {
    store := NewInMemoryStore()
    
    for i := 0; i < b.N; i++ {
        store.FindAll()
    }
}

func BenchmarkTestToDoStoreImpl_FindById(b *testing.B) {
    store := NewInMemoryStore()
    id := 10

    for i := 0; i < b.N; i++ {
        _, _ = store.FindById(id)
    }
}

func BenchmarkTestToDoStoreImpl_Create(b *testing.B) {
	store := NewInMemoryStore()

    for i := 0; i < b.N; i++ {
		task := "Task " + strconv.Itoa(i)
        store.Create(task,"Pending")
    }
}


func BenchmarkTestToDoStoreImpl_Update(b *testing.B) {
	store := NewInMemoryStore()
    id := 10
	task := "Do Something"

    for i := 0; i < b.N; i++ {
       _, _ = store.Update(id, &task, nil)
    }
}


func BenchmarkTestToDoStoreImpl_Delete(b *testing.B) {
	store := NewInMemoryStore()
	id := 9
    for i := 0; i < b.N; i++ {
        store.Delete(id)
    }
}