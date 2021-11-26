package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func partSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func main() {
	part := partSum(3)
	fmt.Println(part(7))
}
