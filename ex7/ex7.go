package main

import (
	"fmt"
	"sync"
)

func write(m map[int]int, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock() //блокирование мютекса на время записи
	m[i] = i
	mutex.Unlock() // разблокирование мютекса
	wg.Done()
}

func main() {
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 0; i < 21; i++ { // конкурентная запись 21 элемента в мапу
		wg.Add(1)
		go write(m, i, &wg, &mutex)
	}

	wg.Wait() // ожидание завершения всех горутин
	fmt.Println(m)
}
