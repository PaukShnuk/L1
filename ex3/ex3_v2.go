package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	wg := sync.WaitGroup{}
	w := sync.RWMutex{}
	arr := []int{2, 4, 6, 8, 10}

	for _, value := range arr {
		wg.Add(1)
		go func(m int) {
			m *= m
			w.Lock() // блокируем мютекс на время запси суммы
			sum += m
			w.Unlock()
			wg.Done()
		}(value)
	}
	wg.Wait() // дожидаемся окончания работы горутин
	fmt.Println(sum)
}
