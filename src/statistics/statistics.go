package main

import "fmt"

func main() {
	x := []int{17, 43, 25, -3, 4, 7}

// 	fmt.Println("The smaller number is:", getSmaller(x))
// 	fmt.Println("The greater number is:", getGreater(x))
// 	fmt.Println("The average is:", getAverage(x))
// 	fmt.Println("Ordered: ", orderArray(x))
}

// func getSmaller(numbers []int) int{
// 	min := numbers[0]

// 	for _, value := range numbers {
// 		if(value < min) {
// 			min = value
// 		}
// 	}

// 	return min
// }

// func getGreater(numbers []int) int {
// 	max := numbers[0]

// 	for i := 0; i < len(numbers); i++ {
// 		if(numbers[i] > max) {
// 			max = numbers[i]
// 		}
// 	}

// 	return max
// }

// func getAverage(numbers []int) int{
// 	var total int
// 	for i := 0; i < len(numbers); i++ {
// 		total += numbers[i]
// 	}

// 	length := len(numbers)
// 	total = total / length
// 	return total
// }

// func orderArray(numbers []int) []int{
// 	change := true
// 	fmt.Println(numbers)
// 	for change {
// 		change = false

// 		for i := 0; i < len(numbers); i++ {
// 			if numbers[i + 1] < numbers[i] {
// 				smallerNumber := numbers[i + 1]
// 				numbers[i + 1] = numbers[i]
// 				numbers[i] = smallerNumber
// 				change = true
// 			}
// 		}

// 	}

// 	return numbers
// }