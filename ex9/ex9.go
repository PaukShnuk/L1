package main

import (
	"fmt"
	"os"
)

func squarer(rCh chan int, sqCh chan int) {
	for value := range rCh {
		sqCh <- value * 2
	}
	close(sqCh)
}

func printer(sqCh chan int, s chan struct{}) {
	for value := range sqCh {
		fmt.Fprintln(os.Stdout, value)
	}
	close(s)
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14}
	readCh := make(chan int)
	squareCh := make(chan int)
	s := make(chan struct{})

	go squarer(readCh, squareCh)
	go printer(squareCh, s)

	for _, value := range arr {
		readCh <- value
	}
	close(readCh)
	<-s

}
