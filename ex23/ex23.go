package main

import "fmt"

func delItem(s []string, index int) {
	s = append(s[:index-1], s[index:]...) // удаление элемента путём комбинирования функции добавления и срезов
	fmt.Println(s)
}

func main() {
	slice := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	delItem(slice, 5)
}
