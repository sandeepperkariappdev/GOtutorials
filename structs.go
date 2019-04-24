package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)

	fmt.Println(sp.age)

	sp1 := s

	fmt.Println(sp1.age)

	sp2 := *sp

	fmt.Println(sp2.age)

}
