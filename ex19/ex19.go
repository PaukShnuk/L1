package main

import (
	"fmt"
)

func revers(s string) {
	s2 := []rune(s) // unicode
	for i := 0; i < len(s2)/2; i++ {
		s2[i], s2[len(s2)-1-i] = s2[len(s2)-1-i], s2[i]
	}
	fmt.Println(string(s2))
}

func main() {
	var s string = "¿Cómo estás?, 世 界! Привет мир"
	revers(s)
}
