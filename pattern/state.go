package main

import (
	"fmt"
)

// State 인터페이스는 상태에 따라 달라지는 동작을 정의
type State interface {
	Operate()
}

// OpenState는 문이 열려 있는 상태를 나타냄
type OpenState struct{}

func (o *OpenState) Operate() {
	fmt.Println("The door closes.")
}

// ClosedState는 문이 닫혀 있는 상태를 나타냄
type ClosedState struct{}

func (c *ClosedState) Operate() {
	fmt.Println("The door opens.")
}

// Door 구조체는 문의 현재 상태를 나타내며, 상태에 따라 동작이 달라짐
type Door struct {
	state State
}

func (d *Door) SetState(s State) {
	d.state = s
}

func (d *Door) Operate() {
	d.state.Operate()
}

func StatePattern() {
	door := &Door{state: &ClosedState{}}

	// 현재 문은 닫혀 있으므로 "The door opens."를 출력
	door.Operate()

	// 문의 상태를 열린 상태로 바꿈
	door.SetState(&OpenState{})

	// 이제 문은 열려 있으므로 "The door closes."를 출력
	door.Operate()
}
