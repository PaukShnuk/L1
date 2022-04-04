package main

import "fmt"

func square(item int, ch chan int) {
	ch <- item * item
}

func main() {
	var sum int
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, value := range arr {
		go square(value, ch) //каждая горутина вычисляет значение квадрата
		sum += <-ch          //складывание результатов горутин
	}
	fmt.Println(sum)

}
