package main

import (
	"fmt"
	"sync"
)

type counter struct {
	c int
}

func increment(c *counter, m *sync.RWMutex, wg *sync.WaitGroup) {
	m.Lock()
	c.c++
	m.Unlock()
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	m := sync.RWMutex{}
	c := new(counter)

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go increment(c, &m, &wg)
	}
	wg.Wait()
	fmt.Println(c.c)
}
