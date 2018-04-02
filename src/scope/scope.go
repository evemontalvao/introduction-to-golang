package main

import "fmt"

func main() {
	x := "Hello"
	fmt.Println(x)

	inside()

	fmt.Println(y)
}

func inside() {
	y := "bye"
}