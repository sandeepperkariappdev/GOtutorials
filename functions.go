package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// <ulitple return in go
func vals() (int, int) {
	return 3, 7
}

func mainFunc() {
	res := plus(1, 2)
	fmt.Println(" 1 + 2", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3", res)

	a, b := vals()
	fmt.Println("a, b", a, b)

	_, c := vals()

	fmt.Println("c", c)
}
