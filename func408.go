package main

import "fmt"

func main() {
	number := 10
	squareNum:= func() int {
		number = number * number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())
	//fmt.Printf("%v\n",squareNum)
	//fmt.Printf("%v",squareNum)
}