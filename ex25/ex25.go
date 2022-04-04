package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	n := time.Now()
	for time.Now().Before(n.Add(t)) { // ждём пока время при вызове функции + время требуемого ожидания не стало меньше, чем текущее
	}
	fmt.Println(time.Now())
}

func main() {
	fmt.Println(time.Now())
	sleep(5 * time.Second)

}
