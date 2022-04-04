package main

import "fmt"

func setBit(current int64, targetBit int64, targetVal int) {
	if targetBit == 0 {
		current = current &^ (1 << (targetBit - 1)) // в случае установления 0 - сброс выбранного бита с помощью И НЕ
	} else {
		current = current | (1 << (targetBit - 1)) // в случае установления 1 - установление бита с помощью ИЛИ
	}
	fmt.Println(current)
}

func main() {
	var current, bit int64
	var value int
	fmt.Println("Введите исходное число, бит который хотите поменять (1-64) и на какое значение хотите поменять (1/0)")
	fmt.Scan(&current, &bit, &value)
	setBit(current, bit, value)
}
