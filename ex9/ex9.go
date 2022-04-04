package main

import (
	"fmt"
	"os"
)

func squarer(rCh chan int, sqCh chan int) {
	for value := range rCh { // получаем данные из первого канала пока он не закроется
		sqCh <- value * 2 // отправляем умноженные данные во второй канал
	}
	close(sqCh) // закрываем второй канал
}

func printer(sqCh chan int, s chan struct{}) {
	for value := range sqCh { // печаем данные из второго канала пока его не закрыли и мы всё не считали
		fmt.Fprintln(os.Stdout, value)
	}
	close(s) // закрываем синхронизирующий канал
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14}
	readCh := make(chan int)
	squareCh := make(chan int)
	s := make(chan struct{})

	go squarer(readCh, squareCh) // запускаем умножитель
	go printer(squareCh, s)      // запускаем принтер

	for _, value := range arr {
		readCh <- value // отправляем данные из массива в первый канал
	}
	close(readCh)
	<-s // ожидаем пока принтер допечатает все значения и закроет канал S

}
