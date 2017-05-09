package main

import "fmt"

func swap(x, y string) (a string,b string) {
	a = y
	b = x
}

func main() {
     	var c = "hello"
	var d string
	d = "world"
	a, b := swap(c, d)
	fmt.Println(a, b)
}
