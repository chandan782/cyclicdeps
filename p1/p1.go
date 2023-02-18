package p1

import (
	"cyclicdeps/p2"
	"fmt"
)

type P1 struct {
}

func New() *P1 {
	return &P1{}
}

func (p P1) HelloFromP1() {
	fmt.Println("Hello from p1")
}

func (p P1) P2UsedInP1() {
	pp2 := p2.New(p)
	pp2.HelloFromP2()
}
