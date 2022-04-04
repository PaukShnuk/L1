package main

import "fmt"

func binarySearch(target int) int {
	var arr = [15]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 20, 21, 24, 26, 29, 33}
	var left, right, mid = 0, len(arr), 0

	for left < right {
		mid = (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return mid
}

func main() {
	var val int
	fmt.Scan(&val)
	fmt.Println("Искомое значение имеет индекс: ", binarySearch(val))
}
