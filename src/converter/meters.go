package main

import "fmt"

func main() {
	fmt.Print("Enter: ")

	var feet float64
	var meters float64 = 0.3048

	fmt.Scanf("%f", &feet)

	feet = feet * meters

	fmt.Println(feet)
}