package main

import (
	"fmt"
	"os"
)

func square(val int, ch chan int) {
	ch <- val * val
}

func main() {
	chInt := make(chan int)
	var arr = []int{2, 4, 6, 8, 10}

	for _, value := range arr {
		go square(value, chInt)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Fprintln(os.Stdout, <-chInt)
	}
}
