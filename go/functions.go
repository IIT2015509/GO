package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var x int
	fmt.Println(add(42, 13))
}
