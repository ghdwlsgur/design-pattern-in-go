package context

import (
	"context"
	"fmt"
	"time"
)

/*
개요
- 고루틴의 실행을 취소하거나 종료하는 데 사용
- 값 전달: 키-값 쌍을 사용하여 고루틴 사이의 값을 전달
- 타임아웃과 데드라인 설정 가능
- 계층적 구조: 부모-자식 관계를 통해 계층적으로 구성 가능
- 고루틴의 수명 주기를 관리함으로써 리소스 누출을 방지
- 네트워크 요청 등에 대해여 타임아웃, 데드라인 설정을 통해 리소스 사용을 제어

장점
- 고루틴에 취소 신호를 전달하여 중단하거나 클린없을 수행가능
- 부모 context로부터 파생된 자식 context를 만들어 설정과 취소 신호를 계층적으로 관리
- 데이터베이스 연결, 네트워크 요청과 같은 프로그램 전체에 걸친 설정 정보를 전달하는 데 사용

단점
- 너무 많은 컨텍스트 값을 전달하거나 잘못된 컨텍스트를 사용하는 경우 복잡성 증가
- 새로운 context를 생성하거나 값들을 전달하고 접근할 때 런타임 오버헤드가 발생할 수 있음
- context의 취소와 관련된 오류 처리 로직이 복잡해질 수 있음

사용 사례
- 요청이 더 이상 유효하지 않을 때 고루틴을 중단하는 데 사용
- 네트워크 요청이나 데이터베이스 쿼리와 같은 작업에 대한 타임아웃을 설정하는 데 유용
- 요청에 대한 메타데이터나 설정 정보를 전달하는 데 사용
- 상위 context에서 하위 고루틴을 취소하거나 관리할 때 사용
*/
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d 취소됨\n", id)
			return
		default:
			fmt.Printf("Worker %d 작업 중\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
