package main

import (
	"fmt"
	"strings"
)

func testUnique(s string) bool {
	s = strings.ToLower(s)  // преобразование всех символов в строке в нижний регистр
	m := make(map[rune]int) // мапка, по которой отслеживаем наличие символа с втроке
	s1 := []rune(s)         //преобразование строки в слайс рун
	for _, value := range s1 {
		if _, ok := m[value]; ok { // проверка на существование записи о букве в мапе
			return false
		}
		m[value]++ // добавление буквы в мапу
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(testUnique(str))
}
