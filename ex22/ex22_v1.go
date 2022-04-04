package main

import (
	"fmt"
	"math/big"
)

func BigAdd(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Add(a1, b1)
	return z
}

func BigSub(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Sub(a1, b1)
	return z
}

func BigMul(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Mul(a1, b1)
	return z
}

func BigDiv(a int64, b int64) big.Int {
	var z big.Int
	a1 := big.NewInt(a)
	b1 := big.NewInt(b)
	z.Div(a1, b1)
	return z
}

func main() {
	var a, b int64 = 11258999068426261, 1125899906842626
	var n int
	fmt.Println("1. Add\n2. Sub\n3. Mul\n4. Div\nEnter №:")
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println(BigAdd(a, b))
	case 2:
		fmt.Println(BigSub(a, b))
	case 3:
		fmt.Println(BigMul(a, b))
	case 4:
		fmt.Println(BigDiv(a, b))
	}
}
