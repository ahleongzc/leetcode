package main

import (
	"fmt"
)

func main() {
	nums := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	countMap := make(map[int]bool)

	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			fmt.Println("duplicate", num)
		}

		countMap[num] = true
	}
}
