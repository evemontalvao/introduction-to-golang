package main

import "fmt"

func main() {
	
	for number := 0; number < 6; number++ {
		if number == 0 {
			fmt.Println("Zero")
		} else if number == 1 {
			fmt.Println("One")
		} else if number == 2 {
			fmt.Println("Two")
		} else if number == 3 {
			fmt.Println("Three")
		} else if number == 4 {
			fmt.Println("Four")
		} else if number == 5 {
			fmt.Println("Five")
		}	
	}
}