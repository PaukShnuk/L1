package main

import "fmt"

func intersection(m1, m2 []int) []int {
	section := make(map[int]int)

	for _, value := range m1 { // проход по первому массиву и запись в мапу кол-ва кажного элемента
		section[value]++
	}

	var m []int //инициализация результирующего слайса

	for _, value := range m2 { // прохождение по второму массиву
		if section[value] >= 1 { // если значение существует в первом массиве, то декрементируем значение в мапе
			section[value]--     // это нужно для нахождения больше одного пересечения одинаковых чисел
			m = append(m, value) // добавляем пересекающийся элемент в результирующий слайс
		}
	}

	return m
}

func main() {
	m1 := []int{2, 4, 6, 8, 2, 2, 2, 2}
	m2 := []int{22, 2, 1, 4, 2, 55, 8, 2}
	fmt.Println(intersection(m1, m2))
}
