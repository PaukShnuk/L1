package main

import "fmt"

type Human struct {
	leftFoot  bool
	rightFoot bool
	leftHand  bool
	rightHand bool
}

type Action struct {
	name     string
	bodyPart Human //встраивание структуры Human как именованое поле
}

func main() {
	a1 := Action{name: "left hook", bodyPart: Human{leftHand: true}}
	a2 := Action{name: "2x jab", bodyPart: Human{leftHand: true, rightHand: true}}
	a3 := Action{"right middle kick", Human{rightFoot: true}}
	fmt.Println(a1, "\n", a2, "\n", a3)
}
