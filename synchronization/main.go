package main

import (
	"fmt"
	"os"
	"pattern/synchronization/semaphore"
	"sync"
	"time"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func semaphoreWithTimeouts() {
	tickets, timeout := 1, 3*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	if err := s.Release(); err != nil {
		panic(err)
	}
}

func semaphoreWithoutTimeouts() {
	tickets, timeout := 0, 0*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	os.Exit(1)
}

func exampleSemaphoreWithTimeout() {
	tickets, timeout := 3, 1*time.Second
	sem := semaphore.New(tickets, timeout)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if err := sem.Acquire(); err != nil {
				fmt.Printf("고루틴 %d: 세마포어 획득 실패: %s\n", id, err)
				return
			}

			fmt.Printf("고루틴: %d: 작업 시작\n", id)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("고루틴 %d: 작업 종료\n", id)

			if err := sem.Release(); err != nil {
				fmt.Printf("고루틴 %d: 세마포어 해제 실패: %s\n", id, err)
			}
		}(i)
	}
	wg.Wait()
}

func main() {

	exampleSemaphoreWithTimeout()
	// semaphoreWithTimeouts()
	// semaphoreWithoutTimeouts()

	// var animal Animal
	//
	// animal = Dog{}
	// fmt.Println(animal.Speak())
	//
	// animal = Cat{}
	// fmt.Println(animal.Speak())
}
