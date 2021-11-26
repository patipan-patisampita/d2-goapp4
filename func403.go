package main

import "fmt"

func area(length, width int) int {
	ar := length * width
	return ar
}
func main() {
	a := area(12,10)
	fmt.Printf("Area of rectangle is: %v\n",a)
	fmt.Println("Area of rectangle is:",a)
}