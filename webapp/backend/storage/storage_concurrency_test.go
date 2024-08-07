package storage

import (
	"sync"
	"testing"
	"todo-webapp/backend/models"
)


func TestConcurrent_todoStore_FindAll(t *testing.T) {

	// Start multiple goroutines to add items concurrently
	var wg sync.WaitGroup
	numGoroutines := 1000
	store := NewInMemoryStore()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			store.FindAll()
		}(i)
	}

	wg.Wait()
}

func TestConcurrent_todoStore_FindById(t *testing.T) {

	// Start multiple goroutines to add items concurrently
	var wg sync.WaitGroup
	numGoroutines := 1000
	store := NewInMemoryStore()
	id := 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			store.FindById(id)
		}(i)
	}

	wg.Wait()
}

func TestConcurrent_todoStore_Create(t *testing.T) {

	// Start multiple goroutines to add items concurrently
	var wg sync.WaitGroup
	numGoroutines := 1000
	store := NewEmptyInMemoryStore()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			store.Create("New item", "Pending")
		}(i)
	}

	wg.Wait()

	items := store.FindAll()

	if hasDuplicatedIds(items) {
		t.Fatalf("storage has repeated ids")
	}

	if len(items) != numGoroutines {
		t.Fatalf("expected %d items, got %d", numGoroutines, len(items))
	}

}

func TestConcurrent_todoStore_Update(t *testing.T) {

	// Start multiple goroutines to add items concurrently
	var wg sync.WaitGroup
	numGoroutines := 1000
	store := NewInMemoryStore()
	id := 1
	task := "Task"

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			store.Update(id, &task, nil)
		}(i)
	}

	wg.Wait()
}


func TestConcurrent_todoStore_Delete(t *testing.T) {

	// Start multiple goroutines to add items concurrently
	var wg sync.WaitGroup
	numGoroutines := 1000
	store := NewInMemoryStore()
	id := 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			store.Delete(id)
		}(i)
	}

	wg.Wait()
}

func hasDuplicatedIds(slice []models.ToDo) bool {
	allKeys := make(map[int]bool)
	for _, item := range slice {
		if _, ok := allKeys[item.Id]; !ok {
			allKeys[item.Id] = true
		}
	}
	return len(allKeys) != len(slice)
}
