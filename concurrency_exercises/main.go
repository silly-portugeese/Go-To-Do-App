package main

import (
	"fmt"
	"sync"
	// "time"
)

// Mutex
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

	for i := 1; i <= 10; i += 2 {
		s.data = i
		fmt.Println("Odd update:", s.data)
	}
}

// Channels
var data int

func odd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		data = i
		fmt.Println("Odd update", data)
		oddCh <- true
		<-evenCh // wait for the even number to be processed before sending the next odd number
	}
}

func even(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-oddCh  // wait for the odd number to be processed
		data = i
		fmt.Println("Even update:", data)
		evenCh <- true
	}
}

func main() {

	// Write a program to simulate a race condition occurring when one goroutine
	// updates a data variable with odd numbers, while another updates the same
	// data variable with even numbers. After each update , attempt to display the
	// data contained in the data variable to screen.
	// [Goroutines][Concurrency][Race Conditions]

	// go run -race concurrency_exercise.go
	// var dt int
	// go func() {

	// 	for i := 1; i <= 10; i += 2 {
	// 		dt = i
	// 		fmt.Println("Odd update:", dt)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }()

	// go func() {
	// 	for i := 2; i <= 10; i += 2 {
	// 		dt = i
	// 		fmt.Println("Even update:", dt)
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
	
	// Case: Wait for one goroutine to finish before starting another
	var wg1 sync.WaitGroup
	safeData := SafeData{}

	wg1.Add(2)

	go func() {
		defer wg1.Done()
		safeData.evenUpdate()
	}()

	go func() {
		defer wg1.Done()
		safeData.oddUpdate()
	}()

	wg1.Wait()

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

	// Case: coordination goroutines to update the value in order 1,2,3,4,5,6,7,8,9,10
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	
	go even(oddCh, evenCh, &wg)
	go odd(oddCh, evenCh, &wg)

	wg.Wait()
	close(evenCh)
	close(oddCh)

	// Output example:
	// Odd update 1
	// Even update: 2
	// Odd update 3
	// Even update: 4
	// Odd update 5
	// Even update: 6
	// Odd update 7
	// Even update: 8
	// Odd update 9
	// Even update: 10

}