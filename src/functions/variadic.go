package main

import "fmt"

func main() {

	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)
}

func sum(nums ...int) {
    total := 0

    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }

    fmt.Println(nums, total)
}

