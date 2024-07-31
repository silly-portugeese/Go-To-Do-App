package main

import (
	"fmt"
	"sync"
)

type SafeData struct {
	mu   sync.Mutex
	data int
}

func (s *SafeData) evenUpdate() {

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 2; i <= 10; i += 2 {
		s.data = i
		fmt.Println("Even update:", s.data)
	}
}

func (s *SafeData) oddUpdate() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 1; i <= 9; i += 2 {
		s.data = i
		fmt.Println("Odd update:", s.data)
	}
}

func main() {

	// Write a program to simulate a race condition occurring when one goroutine
	// updates a data variable with odd numbers, while another updates the same
	// data variable with even numbers. After each update , attempt to display the
	// data contained in the data variable to screen.
	// [Goroutines][Concurrency][Race Conditions]

	// go run -race concurrency_exercise.go
	// var data int
	// go func() {

	// 	for i := 1; i <= 9; i += 2 {
	// 		data = i
	// 		fmt.Println("Odd update:", data)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()

	// go func() {
	// 	for i := 2; i <= 10; i += 2 {
	// 		data = i
	// 		fmt.Println("Even update:", data)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()
	// time.Sleep(time.Second)

	// Output example:
	// Odd update: 3
	// Odd update: 5
	// Even update: 6
	// Even update: 8
	// Odd update: 7
	// Even update: 10
	// Odd update: 9
	// Found 3 data race(s)
	// exit status 66

	// Refactor the program to use channels, mutexes to synchronise all actions.
	var wg sync.WaitGroup
	safeData := SafeData{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		safeData.evenUpdate()
	}()

	go func() {
		defer wg.Done()
		safeData.oddUpdate()
	}()

	wg.Wait()

	// Output example:
	// Odd update: 1
	// Odd update: 3
	// Odd update: 5
	// Odd update: 7
	// Odd update: 9
	// Even update: 2
	// Even update: 4
	// Even update: 6
	// Even update: 8
	// Even update: 10
}
