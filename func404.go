package main

import "fmt"

func rectangle(l int, w int) (area int, parameter int) {
	parameter = 2 * (l + w) //100
	area = l * w            //600
	return
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)

	fmt.Println(p)
	fmt.Println(a)
}
