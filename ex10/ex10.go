package main

import (
	"fmt"
)

func split(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, value := range arr { //прохождение по массиву
		firstNum := int(value / 10)                    //получение первой цифры числа
		m[firstNum*10] = append(m[firstNum*10], value) // получение круглого числа для группы и запись в эту группу значения
	}
	return m
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 4}
	fmt.Println(split(arr))
}
