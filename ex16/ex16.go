package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 { //условие выхода, если массив единичный или нулевой
		return arr
	}
	m, i, j := len(arr)/2, 0, len(arr)-1 //подготовка опорного, левого и правого индексов
	for i <= j {                         //условие схождение границ
		for arr[i] < arr[m] { //пока значение на левой границе меньше опорного
			i++ // сдвиг левой границы вправо
		}
		for arr[j] > arr[m] { //пока значение на правой границе больше опорного
			j-- //
		}
		if i <= j { // если границы не сошлись и знчения надо поменять
			arr[i], arr[j] = arr[j], arr[i] //меняем их местами и двигаем
			i++
			j--
		}
	}
	quickSort(arr[:j+1]) //сортировка левой части сформированного массива
	quickSort(arr[i:])   //сортировка правой части сформированного масива
	return arr
}

func main() {

	var arr = []int{5, -66, 7, -4, 3, -1, 0, 4, 4, 4, 4, 4, -7}
	fmt.Println(arr)
	fmt.Println(quickSort(arr))
}