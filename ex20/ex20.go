package main

import (
	"fmt"
	"strings"
)

func wordReverse(s1 string) {
	s2 := strings.Split(s1, " ") // преобразование строки в слайс строк
	for i := 0; i < len(s2)/2; i++ {
		s2[i], s2[len(s2)-1-i] = s2[len(s2)-1-i], s2[i]
	}
	fmt.Println(strings.Join(s2, " ")) // вывод слайса преобразованного обратно в строку
}

func main() {
	str := "snow dog sun ¿Cómo estás? 世 界"
	fmt.Println(str)
	wordReverse(str)
}
