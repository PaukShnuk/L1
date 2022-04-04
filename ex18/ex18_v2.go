package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	c int64
}

// New Функциия-конструктор
func New() Counter {
	return Counter{}
}

// Increment Используем атомарную операция сложения. Она превосходит мютексы по скорости на небольших объёмах данных
func (c *Counter) Increment() {
	atomic.AddInt64(&c.c, 1)
}

func main() {
	c := New()
	wg := sync.WaitGroup{}

	fmt.Println(c.c)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { // инкрементирование счётчика в конкуренстной среде
			c.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.c)
}
