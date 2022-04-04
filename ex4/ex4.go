package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// функция - work: функции присваивается номер, она читает данные из канала и выводит их.
func work(n int, ch <-chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Fprintln(os.Stdout, "worker's №:", n, "printed:", val)
	}
	wg.Done()
}

func main() {
	var workersNum int
	ch := make(chan int)
	wg := sync.WaitGroup{}
	fmt.Scan(&workersNum)

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go work(i, ch, &wg)
	}

	sigChan := make(chan os.Signal, 1)     //инициализация канала, принимающего сигналы от ОС
	signal.Notify(sigChan, syscall.SIGINT) // уведомление о прерывании в канал

	func() {
		var i int
		for {
			i++
			select {
			case <-sigChan:
				return
			default:
				ch <- i
			}
		}
	}()
	close(sigChan)
	close(ch)
	wg.Wait()
	fmt.Print("stopped")
}
