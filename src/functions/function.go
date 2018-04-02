package main

import "fmt"

func main() {
	x := 79
	y := 45

	result1, result2 := returningFunction(x, y)

	fmt.Println(result2, result1)
}

func returningFunction(x int, y int) (int, int) {
	x = x + 1
	y = y + 1

	return x, y
}

//multiple values
//variadic functions
