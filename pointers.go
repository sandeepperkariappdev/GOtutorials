package main

import "fmt"

// &i  = value at the address i
// *i = address of i
func mainPointers() {
	i := 1
	zeroVal(i)
	fmt.Println("main func execution")
}

func zeroVal(iVal int) {
	//cleariVal = 0
	fmt.Println("ival", iVal)
	//fmt.Println("*ival", *iVal)
	fmt.Println("&ival", &iVal)
	fmt.Println("&ival", *(&iVal))
}

func zeroptr(iptr *int) {
	*iptr = 0
}
