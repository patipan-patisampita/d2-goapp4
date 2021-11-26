package main

import "fmt"

func main() {
	l := 10
	w := 20

	func(){
		var area int
		area = l * w
		fmt.Println(area)
	}()
}