package p2

import "fmt"

type pp1 interface {
	HelloFromP1()
}

type P2 struct {
	P1 pp1
}

func New(p1 pp1) *P2 {
	return &P2{P1: p1}
}

func (p2 *P2) HelloFromP2() {
	fmt.Println("Hello from P2")
}

func (p2 *P2) P1UsedInP2() {
	p2.P1.HelloFromP1()
}
