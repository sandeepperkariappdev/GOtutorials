package main

import "fmt"

func mainRange() {
	nums := []int{2, 3, 4}
	sum := 0
	fmt.Println("nums", nums)
	fmt.Println("sum", sum)
	//for
	for _, num := range nums {
		sum += num
		fmt.Println("sum in  for loop", sum)
		fmt.Println("_ in  for loop")
	}
	fmt.Println("sum after range", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index i , num:", i, num)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}

	fmt.Println("kvs:", kvs)

	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
