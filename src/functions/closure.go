package main

import "fmt"

func evenGenerator() func() int {
  i := 0

  return func() (ret int) {
    ret = i
    i += 2
    return
  }
}

func main() {
  nextEven := evenGenerator()

  fmt.Println(nextEven())
  fmt.Println(nextEven()) 
  fmt.Println(nextEven())
}