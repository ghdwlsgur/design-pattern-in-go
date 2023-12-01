package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- sync.RWMutex는 읽기/쓰기 뮤텍스(Read-Write Mutex)
- 여러 고루틴이 동시에 리소스를 읽을 수 있도록 허용하지만, 쓰기 작업에 대해서는 상호 배제를 사용
- 읽기 잠금(RLock())을 사용하면, 다른 고루틴도 동시에 읽기 잠금을 획득할 수 있지만, 쓰기 잠금(Lock())을 사용하면, 다른 고루틴은 읽기와 쓰기 모두 금지
- 읽기 잠금은 리소스에 대한 동시 읽기를 가능하게 하므로 읽기가 많이 발생하는 상황에서 성능 이점을 제공
*/

var (
	rwMutex sync.RWMutex
	balance int
)

func readBalance(id int, wg *sync.WaitGroup) {
	rwMutex.RLock()
	fmt.Printf("읽기: %d: 잔액: %d\n", id, balance)
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	wg.Done()
}

func writeBalance(value int, wg *sync.WaitGroup) {
	rwMutex.Lock()
	fmt.Printf("쓰기: 잔액 변경: %d\n", value)
	balance = value
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	balance = 100

	wg.Add(3)
	go writeBalance(500, &wg)
	go readBalance(1, &wg)
	go readBalance(2, &wg)
	wg.Wait()

	fmt.Printf("잔액: %d\n", balance)
}
