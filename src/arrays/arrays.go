package main

import "fmt"

func main() {
	var emptyArray [5]int

	filledArray := []int{
		48,96,86,68,
		52,82,47,12,
		21,99,10,45,
		14,43,94,65,
	}

	emptyArray[0] = 1

	fmt.Println(emptyArray)
	fmt.Println(filledArray)
}