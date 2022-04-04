package main

import (
	"context"
	"fmt"
	"time"
)

// worker1: 1 - передача значения в канал, отвечающий на выход из функции
func worker1(q chan struct{}) {
	for {
		select {
		case <-q:
			fmt.Println("Закрытие 1 воркера")
			return
		default:
			fmt.Println("for loop: 1")
		}
		time.Sleep(time.Millisecond)
	}
}

// worker2: 2 - остановка горутины путём закрытия канала
func worker2(ch <-chan int) {
	for range ch {
		fmt.Println("for loop: 2")
	}
	fmt.Println("Закрытие 2 воркера")
}

// worker3: 3 - остановка горутины передачей контекста
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // if cancel() execute
			fmt.Println("Закрытие 3 воркера")
			return
		default:
			fmt.Println("for loop: 3")
		}
		time.Sleep(time.Millisecond)
	}
}

func worker4(b *bool) {
	for {
		if *b {
			fmt.Println("закрытие 4 воркера")

			return
		}
		fmt.Println("for loop: 4")
		time.Sleep(time.Millisecond)
	}
}

func main() {
	//// 1 - передача значения в канал, отвечающий на выход из функции
	quitCh := make(chan struct{})
	go worker1(quitCh)
	time.Sleep(2 * time.Millisecond)
	quitCh <- struct{}{}
	////

	//// 2 - закрытие канала
	ch := make(chan int)
	go worker2(ch)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Nanosecond)
		ch <- i
	}
	close(ch)
	////

	//// 3 - завершение горутины с помощью контекста(схоже с 1 способом)
	ctx, cancel := context.WithCancel(context.Background())
	go worker3(ctx)
	time.Sleep(2 * time.Millisecond)
	cancel()
	////

	//// 4 - завершение передачей переменной
	b := false
	go worker4(&b)
	time.Sleep(2 * time.Millisecond)
	b = true
	////
	time.Sleep(1 * time.Second)
}
