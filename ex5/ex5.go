package main

import (
	"fmt"
	"sync"
	"time"
)

func read(ch <-chan int, wg *sync.WaitGroup) {
	for {
		val, opened := <-ch
		if !opened {
			break
		}
		fmt.Println(val)
	}
	wg.Done()

}

func write(ch chan int, stop <-chan struct{}) {
	var i int
	for {
		select {
		case <-stop: // сигнал завершения работы
			return
		default:
			ch <- i
			i++
		}
	}
}

func main() {
	var s time.Duration
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)

	fmt.Scan(&s) // время N

	chStop := make(chan struct{}) // канал, в который подаётся сигнал о завершении работы

	go write(ch, chStop)
	go read(ch, &wg)
	time.Sleep(s * time.Second)
	chStop <- struct{}{} //отправка пустой структуры в качестве триггера для  завершения записи в канал
	close(ch)

}
