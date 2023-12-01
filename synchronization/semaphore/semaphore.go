package semaphore

import (
	"errors"
	"time"
)

// 에러 정의: 세마포어를 획득하거나 해제할 때 발생할 수 있는 에러 상황 정의
var (
	ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release the semaphore without acquiring it first")
)

// 세마포어 인터페이스 정의, 이 인터페이스는 Acquire와 Release메서드를 가짐
type Interface interface {
	Acquire() error
	Release() error
}

// 세마포어 인터페이스를 구현하는 구조체
type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

// 세마포어 획득 메서드
func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		// 성공적으로 세마포어 획득
		return nil
	case <-time.After(s.timeout):
		// 타임아웃으로 획득 실패
		return ErrNoTickets
	}
}

// 세마포어 해제 메서드
func (s *implementation) Release() error {
	select {
	case _ = <-s.sem:
		// 세마포어 성공적으로 해제
		return nil
	case <-time.After(s.timeout):
		// 타임아웃으로 인해 해제 실패
		return ErrIllegalRelease
	}
}

// 세마포어 생성자 함수, 새로운 세마포어 인스턴스를 생성 및 반환
/*
아래 생성자 함수에서는 인터페이스를 구현하는 implementation 구조체의 인스턴스를 반한하며
반환 타입은 인터페이스로 설정하였는데 이렇게 설정한 이유는 아래와 같습니다.

1. 추상화와 유션성: Interface 타입을 반환함으로써 호출자는 구체적인 구현(struct)에 대해
알 필요가 없게 되며 이는 코드의 결합도를 낮추고 다른 구현으로 쉽게 교체할 수 있도록 해줍니다.
2. 은닉화: 구체적인 구현을 숨기고 인터페이스만을 노출하면서 내부 구현의 변경이 외부에 영향을
미치지 않도록 해주어 코드의 유지보수가 용이해집니다.
*/
func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets), // 지정된 수만큼의 티켓(채널 버퍼 크기)으로 채널 초기화
		timeout: timeout,                      // 타임아웃 값 설정
	}
}
