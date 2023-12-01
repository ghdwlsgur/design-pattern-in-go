package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("입금 중: %d\n", value)
	balance += value
	time.Sleep(1 * time.Second)
	mutex.Unlock()
	wg.Done()
}

/*
- sync.Mutex는 상호 배제를 위한 기본 뮤텍스
- 한 번에 하나의 고루틴만이 리소스에 접근할 수 있도록 보장
- 리소스에 대한 접근을 얻으려면, 고루틴은 Lock()을 호출해야 하고, 작업을 마친 후에는 Unlock()을 호출
- 읽기와 쓰기 모두에 대해 동일한 수준의 접근 제한을 적용합니다. 즉, 리소스를 읽기 위해서도 하나의 고루틴만이 접근할 수 있음
*/
func main() {
	var wg sync.WaitGroup
	balance = 0

	wg.Add(2)
	go deposit(100, &wg)
	go deposit(200, &wg)
	wg.Wait()

	fmt.Printf("잔액: %d\n", balance)
}
