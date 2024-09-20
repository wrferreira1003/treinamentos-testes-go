package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "->", i)
		time.Sleep(time.Second * 1)
	}
	waitGroup.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(4)
	go task("Task 1", &waitGroup)
	go task("Task 2", &waitGroup)
	go task("Task 3", &waitGroup)

	// funcoes anonimas
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Task 4 ->", i)
			time.Sleep(time.Second * 1)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

// Output:
// Task 1 -> 0
// Task 2 -> 0
// Task 3 -> 0
// Task 1 -> 1
// Task 2 -> 1
// Task 3 -> 1
// Task 1 -> 2
// Task 2 -> 2
// Task 3 -> 2
