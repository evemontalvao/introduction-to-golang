package main

import "fmt"

func main() {
	fmt.Print("How is the weather? Enter temperature in Fahrenheit: ")

	var temperature float64

	fmt.Scanf("%f", &temperature)

	temperature = (temperature - 32) * 5/9

	fmt.Println(temperature)
}