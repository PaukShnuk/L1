package main

import "fmt"

type Set struct {
	set []string
}

// AddElem Добавление элементов
func (s *Set) AddElem(str []string) {
	for _, value := range str {
		s.set = append(s.set, value)
	}
}

// DeleteElemSinceStart Удаление одного элемента найденного в начале множества
func (s *Set) DeleteElemSinceStart(str string) (int, error) {
	for index, value := range s.set {
		if value == str {
			s.set = append(s.set[:index], s.set[index+1:]...)
			return 0, nil
		}
	}
	return fmt.Println(str, "is not included")
}

// DeleteElemSinceEnd Удаление одного элемента найденного в конце множества
func (s *Set) DeleteElemSinceEnd(str string) (int, error) {
	for i := len(s.set) - 1; i > 0; i-- {
		if s.set[i] == str {
			s.set = append(s.set[:i], s.set[i+1:]...)
			return 0, nil
		}
	}
	return fmt.Println(str, "is not included")
}

// CheckInclusion Проверка на вхождение
func (s *Set) CheckInclusion(str string) (int, error) {
	for _, value := range s.set {
		if value == str {
			return fmt.Println(str, "is included")
		}
	}
	return fmt.Println(str, "is not included")
}

// CreateSet Создание множества
func CreateSet() Set {
	return Set{[]string{}}
}

func main() {
	var str1 = []string{"²", "dog", "cat", "popug", "kekw"}
	var str2 = []string{"cat", "dog", "cat", "µ", "kekw"}

	c := CreateSet()
	c.AddElem(str1)
	fmt.Println(c)

	c.CheckInclusion("dog")
	c.DeleteElemSinceStart("dog")
	fmt.Println(c)

	c.CheckInclusion("dog")
	c.AddElem(str2)
	c.DeleteElemSinceEnd("cat")
	fmt.Println(c)

}
