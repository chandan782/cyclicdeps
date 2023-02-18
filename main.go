package main

import "cyclicdeps/p1"

func main() {
	pp1 := p1.New()
	pp1.P2UsedInP1()
}
