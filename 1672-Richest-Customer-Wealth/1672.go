package main

import (
	"fmt"
)

func main() {
	nums := [3]int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
