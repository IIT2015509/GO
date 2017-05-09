package main

import "fmt"

func main() {
	i,sum :=1,100 
	
	for sum < 1000 {
		sum += i
		i=i+1
	}
	fmt.Println(sum)
}
