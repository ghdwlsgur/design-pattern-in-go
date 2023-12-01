package main

import (
	"fmt"
	"time"
)

/*
개요
- channel과 context는 Go에서 고루틴 간의 통신 및 동시성 관리를 위해 사용되는 두 가지 메커니즘
- hannel을 사용하여 고루틴 간에 데이터를 전달할 수 있음
- channel은 데이터를 전달하는 동시에 고루틴이 데이터를 보내고 받을 준비가 되었는지 동기화
- channel은 버퍼링을 사용하여 비동기적 또는 동기적으로 동작할 수 있음
- channel을 명시적으로 닫아 더 이상 데이터가 전송되지 않도록 컨트롤 가능

장점
- channel을 사용하면 고루틴 간의 통신이 명확해지고, 코드의 의도를 쉽게 이해할 수 있음
- channel은 동시성 관리를 위한 자연스러운 동기화 메커니즘을 제공
- 공유 자원 접근 시 데이터 레이스 문제를 감소

단점
- 큰 규모의 시스템에서 channel의 사용은 복잡성을 증가
- 채널이 올바르게 닫히지 않을 경우 고루틴 누출이 발생할 수 있음
- 잘못 사용할 경우 데드락을 일으킬 수도 있음

사용 사례
  - 고루틴 간의 데이터 전달
  - 동기화 작업
    (channel은 고루틴이 특정 지점에서 동기화되어야 할 때 사용)
  - 이벤트 또는 메시지 기반 프로그래밍
    (이벤트 또는 메시지를 고루틴 간에 전달할 때 유용)
  - Fan-out, Fan-in 패턴
    (작업을 여러 고루틴에 분산시키고 결과를 수집하는 패턴을 구현할 때 사용)
*/
func worker(id int, done chan bool) {
	fmt.Printf("Worker %d 시작\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d 완료\n", id)
	done <- true
}

func main() {
	done := make(chan bool, 2)
	for i := 1; i <= 2; i++ {
		go worker(i, done)
	}

	<-done
	<-done
}
