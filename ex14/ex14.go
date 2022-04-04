package main

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("this is bool")
	case chan int:
		fmt.Println("This is channel")
	default:
		fmt.Println("this is not a int, string, bool or channel")
	}
}

func main() {
	var a1 int
	var a2 string
	var a3 bool
	var a4 chan int

	checkType(a1)
	checkType(a2)
	checkType(a3)
	checkType(a4)

}
