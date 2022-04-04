package main

import "fmt"

var justString string

// создаём строку с Unicode-символом на месте будущего среза
func createHugeString(l int) string {
	var str string
	for i := 0; i < 99; i++ {
		str += "a"
	}

	str += "ᵺ"

	for i := 0; i < l; i++ {
		str += "a"
	}
	return str
}

/*
		В основе строк, также как и в основе слайсов, лежит ссылка на массив. В отличии от слайсов элементы в
	строке менять нельзя. При создании среза строки или слайса новый массив не создаётся, а созданный слайс ссылается на
	уже существующий массив с указанием len, cap и первого нужного элемента. Таким образом не требуется создавать копии
	одного массива.
*/

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v, len([]byte(v)), cap([]byte(v)))

	/*
			Индексирование строки дает ее байты, а не символы: строка - это просто набор байтов. Это значит, что когда
		мы храним символьное значение в строке, мы храним его байтовое представление. Поэтому, justString = v[:100]
		представляет собой не срез первых 100 символов, а первых 100 байт. В случае, если строка не содержит
		юникод-символов проблем не будет, однако, если в строке содержится такой символ, который занимает уже
		больше 1 байта могут начаться проблемы. В моём случае я использовал символ "ᵺ", занимающий 3 байта и
		стоящий на границе среза.
			Как видно, без преобразований, можно потерять данные, обрезав юникод-символ, если он стоит в конце (как
		показано во втором println'e), либо не вывести однобайтные символы, если юникод-символ стоит в середине.
	*/
	justString = v[:100]
	fmt.Println(justString, len([]byte(justString)), cap([]byte(justString)))

	/*
			Разработчики Го заостряют внимание на юникод-символах и вводят термин "rune" (https://go.dev/blog/strings)
		Для корректной обработки преобразуем строку v в срез рун и выведем первые 100, уже не байт, а рун. Видим, что
		длинна строки а байтах увеличилась на 2, как и положено, а строка выводится корректно.
			Поскольку мы создаем слайс рун у нам создатся новый массив. После, этот массив пареобразуется в строку,
		и создаётся ещё одна копия массива с нужной длинной. Так как первый "рунный" массив не используется, то
		сборщик мусора избавится от него.
			По итогу, в отличии от первого варианта у нас будет инициализировано ещё два массива в памяти,
		но данные сохранятся.
	*/
	c := []rune(v)
	justStringRune := string(c[:100])
	fmt.Println(justStringRune, len([]byte(justStringRune)), cap([]byte(justStringRune)))

}
func main() {
	someFunc()
}
