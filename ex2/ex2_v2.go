package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, val := range arr {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, m*m)
		}(val)
	}

	wg.Wait() // дожидаемя окончания выполнения всех горутин
}
