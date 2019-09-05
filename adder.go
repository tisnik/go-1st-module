package main

import "fmt"

func Add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(0, 0))
	fmt.Println(Add(1, 0))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(-1, 2))
}
